package main

func main() {
	userData := Variables{}
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&userData)
}
