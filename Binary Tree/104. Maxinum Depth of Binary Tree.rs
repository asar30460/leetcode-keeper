

struct Node {
    val: i32,
    left: *Node,
    right: *Node,
}

impl Node {
    fn insert(&self, val: i32) -> Result<_, String> {
        if (self.val == null) {
            self.val = val;
            return Ok()
        }

        let visitor = self;
        let parent: *Node;

        for (self != null) {
            parent = visitor;
            if (self.val < val) {
                visitor = self.right;
            } else  {
                visitor = self.left
            } 
        }

        let new_node = Node {val: val};
        if val > parent.val {
            parent.right = new_node
        } else {
            parent.left = new_node
        }

        
        Ok();
    }
}