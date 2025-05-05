EVALUATE HEALTH 

This is a sample code for "AUTHENTICATION" using ECHO framework of GOLANG 

It uses psql for storing the data.

Folders in the sample code are: 
1: Handlers Folder -- It contains the files which consists of  handler Functions 
2: Models Folder --  It contains struct type for the user model 
3: Repositories -- It contains the files which  consists of operations on the spectific table of the database 
4: Storage -- It contain a single file to connect with database and get the database 
5: Templates -- It contains HTML documents 


main.go -- It contains the all routing of the project. We need to run this file to load the web server.
.env is used to store the details of the database .
