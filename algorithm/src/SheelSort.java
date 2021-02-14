public class SheelSort {
    public void sheelSort(int[] a){
        int b = a.length;
        while(b!=0){
            b=b/2;
            for (int i = 0; i < b; i++) {
                for (int j = i+b; j < a.length; j+=b) {
                    int k = j-b;
                    int temp = a[j];
                    for (; k>=0&&temp<a[k] ; k-=b) {
                        a[k+b] =a[k];
                    }
                    a[k+b]=temp;
                }
            }
        }
    }
}
