public class ListNode
{
    public int val;
    public ListNode next;

    public ListNode(int val = 0, ListNode next = null)
    {
        this.val = val;
        this.next = next;
    }
}

public class Solution
{
    public bool IsPalindrome(ListNode head)
    {
        ListNode middle = head;
        ListNode fast = head;

        while (fast != null)
        {
            middle = middle.next;
            fast = fast.next;

            if (fast.next != null)
            {
                fast = fast.next;
            }
        }

        ListNode tail = null;

        while (middle != null)
        {
            ListNode next = middle.next;
            middle.next = tail;
            tail = middle;
            middle = next;
        }

        while (tail != null)
        {
            if (head.val != tail.val)
            {
                return false;
            }
        }

        return true;
    }

    public class ListNode
    {
        public int val;
        public ListNode next;

        public ListNode(int val = 0, ListNode next = null)
        {
            this.val = val;
            this.next = next;
        }
    }

    public class Solution
    {
        public bool IsPalindrome(ListNode head)
        {
            ListNode middle = head;
            ListNode fast = head;

            while (fast != null)
            {
                middle = middle.next;
                fast = fast.next;

                if (fast.next != null)
                {
                    fast = fast.next;
                }
            }

            ListNode tail = ReverseList(middle);

            while (tail != null)
            {
                if (head.val != tail.val)
                {
                    return false;
                }
                head = head.next;
                tail = tail.next;
            }

            return true;
        }

        private ListNode ReverseList(ListNode head)
        {
            ListNode prev = null;
            ListNode curr = head;

            while (curr != null)
            {
                ListNode nextNode = curr.Next;
                curr.Next = prev;
                prev = curr;
                curr = nextNode;
            }

            return prev;
        }
    }
}
