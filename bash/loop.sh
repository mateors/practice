#!/usr/bin/bash

PS3='Please enter your choice '
select opt in option1 option2 option3 quit; do
     case "$opt" in
       "option1")
            echo "you chose choice 1" ;;
       "option2")
            echo "you chose choice 2" ;;
       "option3")
            echo "you chose choice 3" ;;
       "quit")
            break ;;
       *) echo invalid option ;;
    esac
done
