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
    public string GameResult(ListNode head)
    {
        int oddScore = 0,
            evenScore = 0;

        while (head != null && head.next != null)
        {
            int even = head.val;
            int odd = head.next.val;
            if (even > odd)
            {
                evenScore++;
            }
            else if (odd > even)
            {
                oddScore++;
            }

            head = head.next.next;
        }

        if (oddScore > evenScore)
        {
            return "Odd";
        }
        else if (evenScore > oddScore)
        {
            return "Even";
        }
        else
        {
            return "Tie";
        }
    }
}
