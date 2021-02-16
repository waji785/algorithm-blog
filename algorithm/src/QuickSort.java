public class QuickSort {
    public static void quickSort(int[] a,int start,int end){
        int base = a[start];
        int temp;
        int i =start, j=end;
        do{
            while((a[i]<base)&&(i<end))
                i++;
            while((a[j]>base)&&(j>start))
                j--;
            if(1<=j){
                temp = a[j];
                a[i]=a[j];
                a[j]=temp;
                i++;
                j--;
            }
        }while(i<=j);
        if(start<j)
            quickSort(a,start,j);
        if(end>1)
            quickSort(a,i,end);
    }
}
