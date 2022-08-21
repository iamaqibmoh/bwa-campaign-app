Database Table Requirements
> Users
* id : int
* name : varchar
* email : varchar
* occupation : varchar
* password_hash : varchar
* avatar : varchar
* role : varchar
* token : varchar
* created_at : datetime
* update_at : datetime

> Campaign
* id : int
* user_id : int
* name : varchar
* summary : varchar
* description : text
* goal_amount : int
* current_amount : int
* perks : text
* backer_count : int
* slug : varchar
* created_at : datetime
* update_at : datetime

> Campaign Image
* id : int
* campaign_id : int
* file_name : varchar
* is_primary : boolean (tinyInt)
* created_at : datetime
* update_at : datetime

> Transaction
* id : int
* campaign_id : int
* user_id : int
* amount : int
* status : varchar
* code : varchar
* created_at : datetime
* update_at : datetime
