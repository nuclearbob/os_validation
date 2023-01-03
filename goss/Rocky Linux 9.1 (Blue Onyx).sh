# At time of generation, the yaml file matches AlmaLinux 9.1 (Lime Lynx).yaml and it might be possible to leverage that
GOSSFILE="Rocky Linux 9.1 (Blue Onyx).yaml"
goss --gossfile "$GOSSFILE" autoadd rpm sshd
