//将方法写入jar包并读取调用，目的是混淆隐藏实际调用的方法细节
public class DemoOA {
    public static void main(String[] args) throws Exception {
        Double salary = 2000.00;
        Double money;
        SalaryJarLoader salaryJarLoader = new SalaryJarLoader("F:\\algorithm-blog\\EvlJar\\SalartCaler.jar");
        do {
            money = calSalary(salary,salaryJarLoader);
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
