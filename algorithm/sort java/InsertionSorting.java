public class InsertionSorting {
    public void insertSort(int[] a){
        //写在循环外防止多次初始化提高性能
        int length=a.length;
        int insertNum;
        for (int i = 1; i <length ; i++) {
            insertNum=a[i];
            int j=i-1;
            while (j>=0&&a[j]>insertNum){
                a[j+1]=a[j];
                j--;
            }
            a[j+1]=insertNum;
        }
    }
}
//时间复杂度为0(n)~O(n^2),适合1000个以下的数据排序，不适合1000个以上的排序，插入排序至少进行数组长度次数的遍历。
//插入排序的本质是通过构建小规模的有序序列，将序列不断扩大。