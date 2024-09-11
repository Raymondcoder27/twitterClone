if [ -f /usr/local/bin/.env]; then
    export $(grep -v '^#' /usr/local/bin/.env | xargs)
fi

#Start the backend application
