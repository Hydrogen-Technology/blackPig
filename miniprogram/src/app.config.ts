export default defineAppConfig({
  pages: [
    'pages/index/index',
    'pages/user/index',
    'pages/reserve/index',
    'pages/record/index',
  ],
  window: {
    backgroundTextStyle: 'light',
    navigationBarBackgroundColor: '#fff',
    navigationBarTitleText: 'WeChat',
    navigationBarTextStyle: 'black'
  },
  entryPagePath: 'pages/index/index',
  tabBar: {
    color: '#000',
    selectedColor: '#6190E8',
    backgroundColor: '#fff',
    borderStyle: 'black',
    list: [
      {
        pagePath: 'pages/index/index',
        iconPath: 'assets/images/home.png',
        selectedIconPath: 'assets/images/home_active.png',
        text: '首页', 
      },
      {
        pagePath: 'pages/reserve/index',
        iconPath: 'assets/images/goods.png',
        selectedIconPath: 'assets/images/goods_active.png',
        text: '预约',
      },
      {
        pagePath: 'pages/record/index',
        iconPath: 'assets/images/collect.png',
        selectedIconPath: 'assets/images/collect_active.png',
        text: '记录',
      },
      {
        pagePath: 'pages/user/index',
        iconPath: 'assets/images/user.png',
        selectedIconPath: 'assets/images/user_active.png',
        text: '我的',
      }
    ]
  }
})
