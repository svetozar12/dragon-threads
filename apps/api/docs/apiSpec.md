# BACKEND

## Resources

- Users
- Auth(oauth with google)
- Posts
- PostComments
- Subdragon(works like subreddit)
- Search
- UserInteration
- Rate limitting and Throttling
- ModerationTools
- Voting

## API Specification:

Users:

    Create User: POST /api/v1/users
    Get User by ID: GET /api/v1/users/{user_id}
    Update User: PUT /api/v1/users/{user_id}
    Delete User: DELETE /api/v1/users/{user_id}
    List Users: GET /api/v1/users

Authentication (OAuth with Google):

    Initiate Google OAuth: POST /api/v1/auth/google
    Exchange Google OAuth Code for Access Token: POST /api/v1/auth/google/token
    Logout User: POST /api/v1/auth/logout

Posts:

    Create Post: POST /api/v1/posts
    Get Post by ID: GET /api/v1/posts/{post_id}
    Update Post: PUT /api/v1/posts/{post_id}
    Delete Post: DELETE /api/v1/posts/{post_id}
    List Posts: GET /api/v1/posts

Post Comments:

    Create Comment: POST /api/v1/posts/{post_id}/comments
    Get Comment by ID: GET /api/v1/comments/{comment_id}
    Update Comment: PUT /api/v1/comments/{comment_id}
    Delete Comment: DELETE /api/v1/comments/{comment_id}
    List Comments for Post: GET /api/v1/posts/{post_id}/comments

Subdragons (Subreddits):

    Create Subdragon: POST /api/v1/subdragons
    Get Subdragon by ID: GET /api/v1/subdragons/{subdragon_id}
    Update Subdragon: PUT /api/v1/subdragons/{subdragon_id}
    Delete Subdragon: DELETE /api/v1/subdragons/{subdragon_id}
    List Subdragons: GET /api/v1/subdragons

Voting:

    Upvote Post/Comment: POST /api/v1/posts/{post_id}/upvote
    Downvote Post/Comment: POST /api/v1/posts/{post_id}/downvote

Search Functionality:

    Search Posts: GET /api/v1/search/posts?q={query}
    Search Users: GET /api/v1/search/users?q={query}

Moderation Tools:

    Remove Post/Comment: DELETE /api/v1/moderation/posts/{post_id}
    Ban User: POST /api/v1/moderation/ban/{user_id}
    Set Community Guidelines: PUT /api/v1/subdragons/{subdragon_id}/guidelines

User Interactions:

    Follow User: POST /api/v1/users/{user_id}/follow
    Unfollow User: POST /api/v1/users/{user_id}/unfollow
    Add Friend: POST /api/v1/users/{user_id}/add-friend
    Block User: POST /api/v1/users/{user_id}/block

Notifications:

    List Notifications: GET /api/v1/notifications
    Mark Notification as Read: PUT /api/v1/notifications/{notification_id}/read

Reporting and Flagging:

    Report Post/Comment: POST /api/v1/reports
    Handle Reported Content: PUT /api/v1/moderation/handle-report/{report_id}
