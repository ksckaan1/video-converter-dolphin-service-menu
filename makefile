build:
	go build -o dolphinvc .
install:
	echo "Starting installation..."
	sudo cp ./dolphinvc /usr/bin
	sudo cp ./videoconv.desktop $HOME/.local/share/kservices5/ServiceMenus/
	echo "Installed Successfully"
uninstall:
	echo "Starting Uninstallation..."
	sudo rm -rf /usr/bin/dolphinvc
	sudo rm -rf $HOME/.local/share/kservices5/ServiceMenus/videoconv.desktop
	echo "Uninstalled Successfully"
buildnreinstall: build uninstall install
