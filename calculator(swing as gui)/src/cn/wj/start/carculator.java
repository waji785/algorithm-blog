package cn.wj.start;

import cn.wj.util.Const;

import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class carculator extends  JFrame implements ActionListener {
    //北面控件
    private JPanel jpNorth = new JPanel();
    private JTextField inputText = new JTextField();
    private JButton cBtn = new JButton("C");
    //中间的控件
    private  JPanel jpCenter = new JPanel();
    //初始化
    public carculator()throws HeadlessException {
        this.init();
        this.addNorthComponent();
        this.addCenterButton();
    }
    //绘制窗体
    public void init(){
        this.setTitle(Const.title);
        this.setSize(Const.frameWidth,Const.frameheight);
        this.setLayout(new BorderLayout());
        this.setLocation(Const.frameX,Const.frameY);
        this.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        //GUI是否可拉伸
        this.setResizable(false);
    }
    //添加北面控件
    public  void addNorthComponent(){
        //定义输入框
        this.inputText.setPreferredSize(new Dimension(230,30));
        jpNorth.add(inputText);
        //定义控件颜色
        this.cBtn.setForeground(Color.black);
        jpNorth.add(cBtn);
        //添加控件
        this.add(jpNorth,BorderLayout.NORTH);
        cBtn.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                inputText.setText("");
            }
        });
    }
    //添加中间16个按钮
    public void addCenterButton(){
        //定义按钮名字
        String btnText = "123+456-789*0.=/";
        //正则筛选运算符
        String regex = "[\\+\\-*/.=]";
        //添加按键
        this.jpCenter.setLayout(new GridLayout(4,4));
        //加入名字
        for(int i=0;i<16;i++){
            String temp =btnText.substring(i,i+1);
            JButton btn = new JButton();
            btn.setText(temp);
            if(temp.matches(regex)) {
                btn.setFont(new Font("粗体",Font.BOLD,32));
                btn.setForeground(Color.black);
            }
            btn.addActionListener(this);
            jpCenter.add(btn);
        }
        this.add(jpCenter,BorderLayout.CENTER);
    }
    public static void main(String[] args){
        carculator carculator = new carculator();
        carculator.setVisible(true);
    }
    private String firstInput = null;
    private String operator = null;
    @Override
    //通过拼接实现输入数字
    public void actionPerformed(ActionEvent e) {
        String clickStr = e.getActionCommand();
        if(".0123456789".indexOf(clickStr)!=-1){
            this.inputText.setText(inputText.getText()+clickStr);
            this.inputText.setHorizontalAlignment((JTextField.RIGHT));
        }else if(clickStr.matches("[\\+\\-*/.]{1}")){
            //如果输入的是运算符，那么记录输入的值和运算符然后清空,
            operator = clickStr;
            firstInput = this.inputText.getText();
            this.inputText.setText("");
        }else if(clickStr.equals("=")) {
            Double a = Double.valueOf(firstInput);
            Double b = Double.valueOf(this.inputText.getText());
            Double result = null;
            switch(operator){
                case"+":
                    result =a+b;
                    break;
                case"-":
                    result =a-b;
                    break;
                case"*":
                    result =a*b;
                    break;
                case"/":
                    if(b!=0){
                        result=a/b;
                    };
                    break;
            }
            this.inputText.setText(result.toString());
        }
    }
}
