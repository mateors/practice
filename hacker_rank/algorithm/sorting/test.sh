 APP_PORT="8082"
	#sudo touch /etc/nginx/sites-available/romdevelop.com.conf
 	sudo -u www-data cat > /etc/nginx/sites-available/romdevelop.com.conf << EOL
		server {
			listen 80;
			server_name romdevelop.com www.romdevelop.com;

			root /var/www/romdevelop.com/public_html;

			# Add index.php to the list if you are using PHP
			index index.html;

			location / {
            			proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
		            	proxy_set_header X-Forwarded-Proto \$scheme;
            			proxy_set_header X-Real-IP \$remote_addr;
		            	proxy_set_header Host \$host;
            			proxy_pass  http://127.0.0.1:${APP_PORT};
			}
		}
EOL

if [ $? -eq 0 ]; then
 echo "Virtual hosting created successfully"
fi

sudo chown mostain:www-data /etc/nginx/sites-available/romdevelop.com.conf

if [ $? -eq 0 ]; then
 echo "file permission set successfully"
fi
