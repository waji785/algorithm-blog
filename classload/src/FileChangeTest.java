import java.io.*;

public class FileChangeTest {
    public static void main(String[] args) throws IOException {
        File file = new File("F:\\algorithm-blog\\EvlJar\\out\\production\\EvlJar\\com\\aa\\SalartCaler.class");
        try {
            FileInputStream fis = new FileInputStream(file);
            File targetFile = new File("F:\\algorithm-blog\\EvlJar\\out\\production\\EvlJar\\com\\aa\\SalartCaler.class");
            FileOutputStream fos = new FileOutputStream(targetFile);

            int code =1;
            fos.write(code);
            while ((code = fis.read())!=-1){
                fos.write(code);
            }
            fis.close();
            fos.close();

        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }

    }
}
