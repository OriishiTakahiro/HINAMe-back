SELECT 
  c.board_id, c.id, c.parent_id, c.title, c.author, c.body, c.updated_at 
FROM 
  hiname.comments c 
ORDER BY c.updated_at DESC;
  
SELECT 
  c.id, c.parent_id, c.title, c.author, c.body, c.updated_at 
FROM 
  hiname.comments c 
INNER JOIN 
  hiname.boards b ON c.board_id = b.id 
WHERE 
  b.shelter_id = 10 
ORDER BY c.updated_at DESC;
