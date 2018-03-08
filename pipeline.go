package bitpipe

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/fsouza/go-dockerclient"
)

type Pipeline struct {
	Image        string
	Commands     []string
	WorkDir      string
	OutputStream io.Writer
	Bind         string

	client    *docker.Client
	container *docker.Container
}

func (p *Pipeline) Run() (err error) {

	if err := p.pullImage(); err != nil {
		return err
	}

	if err := p.createContainer(); err != nil {
		return err
	}
	defer p.removeContainer()

	if err := p.startContainer(); err != nil {
		return err
	}
	defer p.stopContainer()

	cmd := []string{"/bin/sh"}
	// runs properly. Using bash does not seem like an elegant solution,
	// but this is the best so far.
	de := docker.CreateExecOptions{
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Tty:          false,
		Cmd:          cmd,
		Container:    p.container.ID,
	}
	log.Debug("CreateExec")
	dExec, err := p.client.CreateExec(de)
	if err != nil {
		log.Debug("CreateExec Error: %s", err)
		return err
	}
	log.Debug("Created Exec")
	execId := dExec.ID

	pr, pw := io.Pipe()
	var errBuffer bytes.Buffer
	mw := io.MultiWriter(&errBuffer, os.Stderr)
	opts := docker.StartExecOptions{
		OutputStream: os.Stdout,
		ErrorStream:  mw,
		InputStream:  pr,
		RawTerminal:  false,
	}

	log.Debug("StartExec")
	cw, err := p.client.StartExecNonBlocking(execId, opts)
	if err != nil {
		log.Debug("CreateExec Error: %s", err)
		return err
	}

	log.Debug("started")

	for _, command := range p.Commands {
		//if command doesnt success return earlier
		io.WriteString(pw, command+" || exit 1 \n")
	}
	io.WriteString(pw, "exit\n")

	defer pw.Close()

	if err := cw.Wait(); err != nil {
		return err
	}

	inspectResult, err := p.client.InspectExec(execId)
	if err != nil {
		return err
	}

	if inspectResult.ExitCode != 0 {
		return errors.New(errBuffer.String())
	}

	return nil
}

func (p *Pipeline) pullImage() error {
	var err error
	if p.client, err = docker.NewClientFromEnv(); err != nil {
		return err
	}

	log.Debug("Created client")

	//Pull image from Registry, if not present
	//imageName := "ubuntu:latest"

	repo, tag := img2RepoandTag(p.Image)
	return p.client.PullImage(docker.PullImageOptions{
		Repository:   repo,
		Tag:          tag,
		OutputStream: os.Stdout,
	}, docker.AuthConfiguration{})
}

func (p *Pipeline) createContainer() error {
	config := docker.Config{
		Image: p.Image,
		// Cmd:          []string{"/bin/sh"},
		WorkingDir: p.WorkDir,

		OpenStdin:    true,
		StdinOnce:    true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
	}
	var err error
	opts2 := docker.CreateContainerOptions{Config: &config}
	opts2.HostConfig = &docker.HostConfig{
		Binds: []string{
			p.Bind,
		},
	}
	p.container, err = p.client.CreateContainer(opts2)
	return err
}

func (p *Pipeline) startContainer() error {
	return p.client.StartContainer(p.container.ID, &docker.HostConfig{})
}

func (p *Pipeline) stopContainer() error {
	return p.client.StopContainer(p.container.ID, 0)
}
func (p *Pipeline) removeContainer() error {
	return p.client.RemoveContainer(docker.RemoveContainerOptions{ID: p.container.ID})
}

func img2RepoandTag(img string) (string, string) {
	s := strings.Split(img, ":")
	if len(s) == 1 {
		return s[0], ""
	}
	return s[0], s[1]
}
