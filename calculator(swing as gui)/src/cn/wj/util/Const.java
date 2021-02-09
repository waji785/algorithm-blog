package cn.wj.util;

import java.awt.*;

public class Const {
    //获取屏幕宽和高
    public static final int screenWidth = Toolkit.getDefaultToolkit().getScreenSize().width;
    public static final int screenheight = Toolkit.getDefaultToolkit().getScreenSize().height;
    //定义计算器窗口大小
    public static final int frameWidth = 300;
    public static final int frameheight = 300;
    //坐标点位置
    public static int frameX = (screenWidth-frameWidth)/2;
    public static int frameY = (screenheight-frameheight)/2;
    //名字
    public static final String title = "carculatot";
}
