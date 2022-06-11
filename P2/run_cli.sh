echo host name = $1
numb=""
numb=$(ls -l /home | grep -a "rootfs" | wc -l) #gets the number of containers + 1(because of the tar file)
#echo container : $numb
mkdir /home/rootfs$numb #makes containers file in system
sudo tar -xf /home/rootfs.tar.gz -C /home/rootfs$numb #extract the tarball file for new container

#sudo cp -R /home/rootfs/* /home/rootfs$numb/

if [$2 == ""]
then
echo NO RAM Limit
sudo go run My_cli.go run /bin/bash $1 $numb # no ram limit
else
echo RAM Limit = $2
systemd-run --scope -p MemoryLimit=$2M go run My_cli.go run /bin/bash $1 $numb
fi

rm -r /home/rootfs$numb #removes the container directory 