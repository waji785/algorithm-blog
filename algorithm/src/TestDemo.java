import java.util.Arrays;

public class TestDemo {
    public static void main(String[] args) {
        //插入排序测试
        InsertionSorting insertionSorting = new InsertionSorting();
        int[] b = new int[]{2,4,3,7,322,13,2,0};
        insertionSorting.insertSort(b);
        System.out.println(Arrays.toString(b));

        //希尔排序测试
        SheelSort sheelSort = new SheelSort();
        int[] c = new int[]{41,1421,315,7,22,215,2,0,1241,233,12};
        sheelSort.sheelSort(c);
        System.out.println(Arrays.toString(c));
    }
}
