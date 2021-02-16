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

        //简单排序
        SelectSort selectSort = new SelectSort();
        int[] d = new int[]{12,3153,14214,1231,313,131,14143,25};
        selectSort.selectSort(d);
        System.out.println(Arrays.toString(d));

        //冒泡排序
        BubbleSort bubbleSort = new BubbleSort();
        int[] e =new int[]{121,314,213,252,12141,213457,8534,7726,2,5252,626,25,2,75474,73};
        bubbleSort.bubbleSort(e);
        System.out.println(Arrays.toString(e));
    }
}
