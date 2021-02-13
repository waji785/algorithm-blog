import java.net.URL;
import java.net.URLClassLoader;
//将方法写入jar包并读取调用，目的是混淆隐藏实际调用的方法细节
public class DemoOA {
    public static void main(String[] args) throws Exception {
        Double salary = 2000.00;
        Double money;
//        URL jarPath = new URL("file:F:\\algorithm-blog\\EvlJar\\ide-eval-resetter-2.1.13.jar");
//        URLClassLoader urlClassLoader = new URLClassLoader(new URL[] {jarPath});
        SalaryClassLoader salaryClassLoader = new SalaryClassLoader("F:\\algorithm-blog\\EvlJar\\");
        do {
            money = calSalary(salary,salaryClassLoader);
            System.out.println("实际到手工资" + money);
            Thread.sleep(1000);
        } while (true);
        }
        private static Double calSalary(Double salary,ClassLoader ClassLoader) throws Exception{
            Class<?> clazz = ClassLoader.loadClass("SalartCaler");
            Object object = clazz.newInstance();
            return (Double)clazz.getMethod("cal",Double.class).invoke(object,salary);
    }
}
