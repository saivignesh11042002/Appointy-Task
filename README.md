# Appointy-Task
This is the work of S.Sai Vignesh. This Backend API was developed as a part of the task assigned by Appointy for internship selection.

All the key and value in the JSON FORMAT must be given only in lower-case

Operation-1: Creating a User

When a "POST" request for the URL "/users" comes in the function user_handle gets fired up. Inside if there is a seperate if clause that deals with the "POST" request. The details of the user to be created is given in the Body of the JSON request. IT IS NECESSARY THAT THE "KEY" IN THE KEY VALUE PAIRS IN THE JSON REQUEST BE GIVEN IN LOWER-CASE AND THE KEYS BE THE EXACT SAME KEYS GIVEN IN THE IMAGE

 

Operation-2: Getting user based on user ID

When a get request for the URL "/users" comes in the function user_handle gets fired up. Inside if there is a seperate if clause that deals with the "GET" request. The query is passed in the URL which is extracted and the user details corresponding to that ID is fetched.

Operation-3: Creating a Post

When a "POST" request for the URL "/posts" comes in the function post_handle gets fired up. Inside if there is a seperate if clause that deals with the "POST" request. The details of the user to be created is given in the Body of the JSON request. The "uid" is IT IS NECESSARY THAT THE KEY VALUE PAIRS IN THE JSON REQUEST BE GIVEN IN LOWER-CASE AND THE KEYS BE THE EXACT SAME KEYS GIVEN IN THE IMAGE

 

Operation-4: Getting a Post using the ID

When a get request for the URL "/users" comes in the function user_handle gets fired up. Inside if there is a seperate if clause that deals with the "GET" request. The query is passed in the URL which is extracted and the post details corresponding to that ID is fetched.

Operation-4: Getting a Post of a user using the USER_ID

When a get request for the URL "posts/users" comes in the function user_posts_handle gets fired up. Inside if there is a seperate if clause that deals with the "GET" request. The query is passed in the URL which is extracted and the posts details of all the posts of the user corresponding to that ID is fetched.

