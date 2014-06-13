#DEIS:
##Architecture Description
###Controller's Responsabilities:
	- Authenticating and authorizing clients;
	- Processing client API calls;
	- Mananging a cluster with nodes to host services;
	- Managing containers that perform work for applications;
	- Managing proxies that route traffic to containers;
	- Managing users, keys and other base configuration.
###Controller's Stack:
	- Django API Server for handling API Calls;
	- Celery, backed by Redis, for dispatching tasks.
	
###Database's responsability:
	- Store durable platform state.
###Database system: 
	- PostgreSQL.
	
###Cache's responsabilies:
	- Store work queue data for CELERY;
	- Cache sessions and synchronize locks for Django;
	- Store recent log data for the Controller.
Cache uses Redis.

###Builder's Responsabilities:
	- Receives incoming git push request over SSH;
	- Authenticates the user via SSH key fingerprint;
	- Authorizes the user's access to write to the Git repository;
	- Builds a new Docker image from the updated git repository;
	- Adds the latest Config to the resulting Docker image;
	- Pushes the new Docker image to the platform's Registry;
	- Creates a new Release on the Controller.

###Registry's Responsabilities:
	- Hosts Docker images on behalf of the platform;
Registry stores the Images data on storage services like AS3 or OSS.

###Log Server 
	- uses rsyslog to aggregate log data across the platform. (data can then be queried by the controller).

###Router
	- uses Nginx to rout traffic to application containers.