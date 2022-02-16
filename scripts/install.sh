#!/bin/bash

installPath=$1
instalPath=""

if [ "$installPath" != "" ]; then
    instalPath=$installPath
else
    instalPath=/usr/local/bin
fi

UNAME=$(uname)
ARCH=$(uname -m)

rmOldFiles() {
    if [ -f $instalPath/instal ]; then
        sudo rm -rf $instalPath/instal*
    fi
}

v=$(curl --silent "https://api.github.com/repos/abdfnx/instal/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

releases_api_url=https://github.com/abdfnx/instal/releases/download

successInstall() {
    echo "🙏 Thanks for installing Instal! If this is your first time using the CLI, be sure to run `instal --help` first."
}

mainCheck() {
    echo "Installing instal version $v"
    name=""

    if [ "$UNAME" == "Linux" ]; then
        if [ $ARCH = "x86_64" ]; then
            name="instal_linux_${v}_amd64"
        elif [ $ARCH = "i686" ]; then
            name="instal_linux_${v}_386"
        elif [ $ARCH = "i386" ]; then
            name="instal_linux_${v}_386"
        elif [ $ARCH = "arm64" ]; then
            name="instal_linux_${v}_arm64"
        elif [ $ARCH = "arm" ]; then
            name="instal_linux_${v}_arm"
        fi

        instalURL=$releases_api_url/$v/$name.zip

        wget $instalURL
        sudo chmod 755 $name.zip
        unzip $name.zip
        rm $name.zip

        # instal
        sudo mv $name/bin/instal $instalPath

        rm -rf $name

    elif [ "$UNAME" == "Darwin" ]; then
        if [ $ARCH = "x86_64" ]; then
            name="instal_macos_${v}_amd64"
        elif [ $ARCH = "arm64" ]; then
            name="instal_macos_${v}_arm64"
        fi

        instalURL=$releases_api_url/$v/$name.zip

        wget $instalURL
        sudo chmod 755 $name.zip
        unzip $name.zip
        rm $name.zip

        # instal
        sudo mv $name/bin/instal $instalPath

        rm -rf $name

    elif [ "$UNAME" == "FreeBSD" ]; then
        if [ $ARCH = "x86_64" ]; then
            name="instal_freebsd_${v}_amd64"
        elif [ $ARCH = "i386" ]; then
            name="instal_freebsd_${v}_386"
        elif [ $ARCH = "i686" ]; then
            name="instal_freebsd_${v}_386"
        elif [ $ARCH = "arm64" ]; then
            name="instal_freebsd_${v}_arm64"
        elif [ $ARCH = "arm" ]; then
            name="instal_freebsd_${v}_arm"
        fi

        instalURL=$releases_api_url/$v/$name.zip

        wget $instalURL
        sudo chmod 755 $name.zip
        unzip $name.zip
        rm $name.zip

        # instal
        sudo mv $name/bin/instal $instalPath

        rm -rf $name
    fi

    # chmod
    sudo chmod 755 $instalPath/instal
}

rmOldFiles
mainCheck

if [ -x "$(command -v instal)" ]; then
    successInstall
else
    echo "Download failed 😔"
    echo "Please try again."
fi
