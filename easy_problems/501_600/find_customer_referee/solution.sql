/* Write your PL/SQL query statement below */
select name from customer where nvl(referee_id,0) <> 2
