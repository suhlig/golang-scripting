# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|  
  config.vm.define "trusty" do |trusty|
    trusty.vm.box = "ubuntu/trusty64"
  end

  config.vm.define "xenial" do |xenial|
    xenial.vm.box = "ubuntu/xenial64"
  end

  config.vm.provision "shell", inline: <<~SHELL
    add-apt-repository ppa:gophers/archive
    apt-get update
    apt-get -y install golang-1.9
    echo 'export PATH=/usr/lib/go-1.9/bin:\$PATH' > /etc/profile.d/go-1.9.sh
  SHELL
end
