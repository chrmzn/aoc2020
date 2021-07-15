# First check Homebrew is installed, if not, install
if [ -z `which brew` ] 
then
    echo "Homebrew not installed ❌🍺"
    curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh
else
    echo "Homebrew installed 🍺"
fi

echo "Updating Homebrew! 🍻"
brew update

# Then check if pyenv is installed, if not, install
if [ -z `which pyenv` ]
then
    echo "Install pyenv and build environment 🐍"
    brew install openssl readline sqlite3 xz zlib pyenv
    if [ $SHELL == "/bin/zsh" ]
    then
        echo "ZSH Shell"
        echo 'eval "$(pyenv init --path)"' >> ~/.zprofile
        echo 'eval "$(pyenv init -)"' >> ~/.zshrc
        eval "$(pyenv init -)"
    elif [ $SHELL == "/bin/bash" ]
    then
        echo "Bash Shell"
        echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.profile
        echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.profile
        echo 'eval "$(pyenv init --path)"' >> ~/.profile
        echo 'eval "$(pyenv init -)"' >> ~/.bashrc
        eval "$(pyenv init -)"
    else
        echo "Cannot determine default shell 🐌 :( Please configure manually (https://github.com/pyenv/pyenv#installation) and re-run this script"
        exit 1
    fi
else
    echo "pyenv installed 🐍"
fi

# Then check for version 3.9.0 of python in Pyenv, if not, install
PYTHON_VERSION_390=`pyenv versions | grep 3.9.0`
if [ -z $PYTHON_VERSION_390 ]
then
    echo "Install Python 3.9.0 🐍"
    pyenv install 3.9.0
else
    echo "Python 3.9.0 already installed 🐍"
fi

# Then check for poetry, if not install
if [ -z `which poetry` ]
then
    echo "Install Poetry 🎼"
    curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | `pyenv which python3.9` -
else
    echo "Poetry already installed 🎼"
fi
