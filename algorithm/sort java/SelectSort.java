public class SelectSort {
    public void selectSort(int[] a){
        int length = a.length;
        for (int i = 0; i < length; i++) {
            int key = a[i];
            int position = i;
            for (int j = i+1; j < length; j++) {
                if(a[j]<key){
                    key = a[j];
                    position = j;
                }

            }
            a[position]=a[i];
            a[i]=key;
        }
    }
}
//每一次都找出最小的数放在最前面
