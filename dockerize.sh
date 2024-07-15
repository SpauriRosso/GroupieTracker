#!/bin/bash

read -p "Enter your choice. Install(y) / Uninstall(n) / Cancel(c):- " choice

if [ "$choice" = "y" ]; then
  if [ -z "$(docker images -q groupie-image:latest 2> /dev/null)" ]; then
    echo "groupie-image does not exist"
    echo "Building image"
    docker build -t groupie-image .
    echo "Starting Container"
    docker run -d --name groupie -p 5826:5826 groupie-image
    echo "Installation completed! Running on port 5826"
  else
    echo "groupie-image:latest already exists"
    read -p "Do you want to launch the container (l) or uninstall the image (u)? " image_choice
    if [ "$image_choice" = "l" ]; then
      if [ "$(docker ps -q -f name=groupie)" ]; then
        echo "Container 'groupie' is already running"
      else
        echo "Starting Container"
        docker run -d --name groupie -p 5826:5826 groupie-image
        echo "Running on port 5826"
      fi
    elif [ "$image_choice" = "u" ]; then
      choice="n"
    else
      echo "Invalid choice. Exiting."
      exit 1
    fi
  fi
fi

if [ "$choice" = "n" ]; then
  echo "Uninstalling..."
  if [ "$(docker ps -q -f name=groupie)" ]; then
    echo "Stopping and removing 'groupie' container"
    docker stop groupie
    docker rm groupie
  fi
  if [ "$(docker images -q groupie-image:latest)" ]; then
    echo "Removing 'groupie-image' image"
    docker rmi groupie-image:latest
  fi
  echo "Uninstallation completed"
elif [ "$choice" = "c" ]; then
  echo "Installation cancelled"
elif [ "$choice" != "y" ]; then
  echo "Wrong choice! Please enter y, n, or c."
fi