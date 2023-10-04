Vagrant.configure("2") do |config|
  config.vm.define "ubuntu" do |ubuntu|
    ubuntu.vm.box = "ubuntu/focal64"
    ubuntu.vm.network "forwarded_port", guest: 80, host: 4000
    ubuntu.vm.network "private_network", ip: "192.168.56.10"
    ubuntu.vm.synced_folder ".", "/vagrant_data"
    ubuntu.vm.provision "shell", inline: <<-SHELL
      apt-get update
      curl -fsSL https://get.docker.com -o get-docker.sh
      sh get-docker.sh
      rm get-docker.sh
      apt-get install nano -y
    SHELL
  end

  config.vm.define "centos" do |centos|
    centos.vm.box = "centos/7"
    centos.vm.network "private_network", ip: "192.168.56.11"
  end
end