public class MergeSort {
    public static void mergeSort(int[] arr,int low,int high) {
        int middle=(high+low)/2;
        if(low<high) {
            mergeSort(arr, low, middle);
            mergeSort(arr, middle+1, high);
            merge(arr,low,middle,high);
        }
    }
    public static void merge(int[] arr,int low,int middle, int high) {
        int[] temp = new int[high-low+1];
        int i=low;
        int j=middle+1;
        int index=0;
        while(i<=middle&&j<=high) {
            if(arr[i]<=arr[j]) {
                temp[index]=arr[i];
                i++;
            }else {
                temp[index]=arr[j];
                j++;
            }
            index++;
        }
        while(j<=high) {
            temp[index]=arr[j];
            j++;
            index++;
        }
        while(i<=middle) {
            temp[index]=arr[i];
            i++;
            index++;
        }
        for(int k=0;k<temp.length;k++) {
            arr[k+low]=temp[k];
        }
    }
}