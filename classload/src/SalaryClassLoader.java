import com.sun.xml.internal.ws.util.ByteArrayBuffer;

import java.io.File;
import java.io.FileInputStream;
import java.security.SecureClassLoader;
//定义类加载器
public class SalaryClassLoader extends SecureClassLoader {
    private String classPath;
    //文件路径
    public SalaryClassLoader(String classPath){

        this.classPath = classPath;
    }
//重写方法，将类的后缀改为myclass，并向二进制文件中写入无效数据防止反编译
    @Override
    protected Class<?> findClass(String name) throws ClassNotFoundException {
        String filePath = this.classPath+name.replace(".","\\").concat(".myclass");
        FileInputStream fis;
        ByteArrayBuffer ba = new ByteArrayBuffer();
        int code = 1;
        byte[] b;
        try{
            fis = new FileInputStream(new File(filePath));
            while((code=fis.read())!=-1){
                ba.write(code);
            }
            b = ba.toByteArray();
            return this.defineClass(name,b,0,b.length);
        }catch (Exception e){
            throw new ClassNotFoundException();
        }
    }
}
