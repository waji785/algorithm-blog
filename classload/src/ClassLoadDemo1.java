public class ClassLoadDemo1 {
    public static void main(String[] args) throws Exception{
        //加载器继承关系
        // AppClassLoader <- ExtClassLoader <- BootStrap ClassLoader
        ClassLoader cl1 = ClassLoadDemo1.class.getClassLoader();
        System.out.println("cli>"+cl1);
        System.out.println("parent of cl1>"+cl1.getParent());
        //BootStrap ClassLoader因为是c++写的，本身是jvm的一部分，不是java类所以打不出来；
        System.out.println("grand parent of cl1>"+cl1.getParent().getParent());
        //String,Int等类由BootStrap ClassLoader加载。
        ClassLoader cl2 = String.class.getClassLoader();
        System.out.println("cl2>" +cl2);
        System.out.println(cl1.loadClass("java.util.List").getClass().getClassLoader());

        //java指令可以通过增加-verbose:class -verbose:gc参数在启动时打印出2类加载情况
        //BootStrap ClassLoader,加载java基础类，这个属性不能在java指令中指定，因此它可能不是java写的
        System.out.println("BootStrap ClassLoader加载目录："+System.getProperty("sun.boot.class.path"));
        //Extention ClassLoder加载java_home/ext下的jar包，可通过-D java.ext.dirs另行指定目录
        System.out.println("Extenstion ClassLoader加载目录"+System.getProperty("java.ext.dirs"));
        //AppClassLoader加载CLASSPATH,应用下的jar包，可通过-D java.class.path另行指定目录
        System.out.println("AppClassLoader加载目录"+System.getProperty("java.class.path"));
    }
}
