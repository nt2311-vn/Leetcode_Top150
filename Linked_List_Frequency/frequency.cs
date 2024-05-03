using System;
using System.Collections.Generic;

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
    public ListNode FrequenciesOfElements(ListNode head)
    {
        Dictionary<int, int> freqMap = new Dictionary<int, int>();

        while (head != null)
        {
            if (!freqMap.ContainsKey(head.val))
            {
                freqMap.Add(head.val, 0);
            }

            freqMap[head.val]++;
            head = head.next;
        }

        ListNode dummyNode = new ListNode();
        ListNode curr = dummyNode;

        foreach (var val in freqMap.Values)
        {
            curr.next = new ListNode(val);
            curr = curr.next;
        }

        return dummyNode.next;
    }
}
