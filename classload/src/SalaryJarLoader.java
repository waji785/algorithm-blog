import com.sun.xml.internal.ws.util.ByteArrayBuffer;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.net.URL;
import java.security.SecureClassLoader;
//定义类加载器
public class SalaryJarLoader extends SecureClassLoader {
    private String jarFile;
    //文件路径
    public SalaryJarLoader(String jarFile){

        this.jarFile = jarFile;
    }
    //重写方法，将类的后缀改为myclass，并向二进制文件中写入无效数据防止反编译
    @Override
    protected Class<?> findClass(String name) throws ClassNotFoundException {
        FileInputStream fis;
        String classPath = name.replace(".","/").concat(".myclass");
        InputStream inputStream;
        URL fileUrl;
        ByteArrayBuffer ba = new ByteArrayBuffer();
        int code = 1;
        byte[] b;
        try{
            fileUrl = new URL("jar:file:\\"+this.jarFile+"!/"+classPath);
            inputStream = fileUrl.openStream();;
            while((code=inputStream.read())!=-1){
                ba.write(code);
            }
            b = ba.toByteArray();
            return this.defineClass(name,b,0,b.length);
        }catch (Exception e){
            throw new ClassNotFoundException();
        }
    }
}
//类加载分为加载、连接、初始化三个部分
//加载loading，用户可以参与的阶段，自定义类加载器就是这个过程，将代码变成JVM可识别的代码。
//连接linking，由验证，检查字节码是否符合规范；准备，创建类或接口的的静态变量并赋值；解析，将引用变成直接引用。预分配内存；
//初始化initalization，执行static代码块初始化，给变量赋值等等。

//另一方面，热加载机制存在缺陷，因为其底层的resolve可以不执行连接阶段，因此将一些编译时期可以检查出来的问题延迟到连接阶段。