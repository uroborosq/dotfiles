echo "Welcome to the UroborosQ's update script!"

echo "yay updating..."

yay

echo "flatpak updating..."

flatpak update

current_directory=$(pwd)


echo "git repos updating..."

echo "libadwaita-theme-changer"
cd /home/uroborosq/.local/share/libadwaita-theme-changer/

git pull


cd $current_directory
