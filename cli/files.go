package cli

import (
	"bytes"
	"os"

	"github.com/digitalrebar/provision/client/files"
	"github.com/spf13/cobra"
)

type FileOps struct{ CommonOps }

func (be FileOps) GetIndexes() map[string]string {
	return map[string]string{}
}

func (be FileOps) List(parms map[string]string) (interface{}, error) {
	d, e := session.Files.ListFiles(files.NewListFilesParams(), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be FileOps) Get(path string) (interface{}, error) {
	b := bytes.NewBuffer(nil)
	_, e := session.Files.GetFile(files.NewGetFileParams().WithPath(path), basicAuth, b)
	if e != nil {
		return nil, e
	}
	noPretty = true
	return string(b.Bytes()), nil
}

func (be FileOps) Upload(path string, f *os.File) (interface{}, error) {
	d, e := session.Files.UploadFile(files.NewUploadFileParams().WithPath(path).WithBody(f), basicAuth)
	if e != nil {
		return nil, e
	}
	return d.Payload, nil
}

func (be FileOps) Delete(id string) (interface{}, error) {
	_, e := session.Files.DeleteFile(files.NewDeleteFileParams().WithPath(id), basicAuth)
	if e != nil {
		return nil, e
	}
	return "Good", nil
}

func init() {
	tree := addFileCommands()
	App.AddCommand(tree)
}

func addFileCommands() (res *cobra.Command) {
	singularName := "file"
	name := "files"
	d("Making command tree for %v\n", name)
	res = &cobra.Command{
		Use:   name,
		Short: "Commands to manage files on the provisioner",
	}
	commands := commonOps(&FileOps{CommonOps{Name: name, SingularName: singularName}})
	res.AddCommand(commands...)
	return res
}
