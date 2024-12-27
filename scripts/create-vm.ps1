$env:VAGRANT_CWD="test/vagrant"
vagrant up
vagrant ssh -c "sudo reboot"