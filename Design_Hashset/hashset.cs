public class MyHashSet
{
    private bool[] payload;

    public MyHashSet()
    {
        payload = new bool[1000001];
    }

    public void Add(int key)
    {
        payload[key] = true;
    }

    public void Remove(int key)
    {
        payload[key] = false;
    }

    public bool Contains(int key)
    {
        return payload[key];
    }
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet obj = new MyHashSet();
 * obj.Add(key);
 * obj.Remove(key);
 * bool param_3 = obj.Contains(key);
 */
