GOSSFILE=ubuntu_22_04.yaml
goss --gossfile $GOSSFILE autoadd apt sshd
ln -s "Ubuntu 22.04.1 LTS.yaml" $GOSSFILE