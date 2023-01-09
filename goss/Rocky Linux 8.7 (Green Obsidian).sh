# At time of generation, the yaml file matches AlmaLinux 8.7 (Stone Smilodon) and it might be possible to leverage that
# Might want to add yum when we add other things later
GOSSFILE="Rocky Linux 8.7 (Green Obsidian)"
goss --gossfile "$GOSSFILE" autoadd rpm sshd
