export URL="https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt"
export namedir="$1 at $(date)"

mkdir "$namedir"
mkdir "$namedir/about_me"
mkdir "$namedir/about_me/personal"
mkdir "$namedir/about_me/professional"
mkdir "$namedir/my_friends"
mkdir "$namedir/my_system_info"

echo "$2" > "$namedir/about_me/personal/facebook.txt"
echo "$3" > "$namedir/about_me/professional/linkedin.txt"
curl "$URL" -o "$namedir/my_friends/list_of_my_friends.txt"
uname -a > "$namedir/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "$namedir/my_system_info/internet_connection.txt"

tree "$namedir"
