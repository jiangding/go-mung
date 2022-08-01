import { fetchUserCenter } from '../../services/usercenter/fetchUsercenter';
import Toast from 'tdesign-miniprogram/toast/index';

const menuData = [
  [
    {
      title: '收货地址',
      tit: '',
      url: '',
      type: 'address',
    },
    {
      title: '优惠券',
      tit: '',
      url: '',
      type: 'coupon',
    },
    {
      title: '积分',
      tit: '',
      url: '',
      type: 'point',
    },
  ],
  [
    {
      title: '帮助中心',
      tit: '',
      url: '',
      type: 'help-center',
    },
    {
      title: '客服热线',
      tit: '',
      url: '',
      type: 'service',
      icon: 'service',
    },
  ],
];

const orderTagInfos = [
  {
    title: '待付款',
    iconName: 'wallet',
    orderNum: 0,
    tabType: 5,
    status: 1,
  },
  {
    title: '待发货',
    iconName: 'deliver',
    orderNum: 0,
    tabType: 10,
    status: 1,
  },
  {
    title: '待收货',
    iconName: 'package',
    orderNum: 0,
    tabType: 40,
    status: 1,
  },
  {
    title: '待评价',
    iconName: 'comment',
    orderNum: 0,
    tabType: 60,
    status: 1,
  },
  {
    title: '退款/售后',
    iconName: 'exchang',
    orderNum: 0,
    tabType: 0,
    status: 1,
  },
];

const getDefaultData = () => ({
 //...
});

Page({
  data: {
    showMakePhone: false,
    // 用户信息
    userInfo: {},
    menuData,
    orderTagInfos,
    customerServiceInfo: {},
    currAuthStep: 1,
    showKefu: true,
    versionNo: '',
  },
  onLoad() {
    this.getVersionInfo();
  },

  onShow() {
    this.getTabBar().init();
    this.init();
  },
  onPullDownRefresh() {
    this.init();
  },

  init() {
    //this.fetUseriInfoHandle();

    // 获取存储的用户授权信息
    const userinfo =  wx.getStorageSync('userinfo') || {}
    console.log(userinfo);
    // 判断是否存在已经授权的用户信息
    if (Object.keys(userinfo).length == 0) {
      this.setData({
        userInfo: {},
        currAuthStep:1,
      })
    } else {
      // 设置data
      this.setData({
        userInfo: {
          nickName: JSON.parse(userinfo).nickName,
          avatarUrl:JSON.parse(userinfo).avatarUrl,
        },
        currAuthStep:2,
      })
    }
    
    wx.stopPullDownRefresh();
  },

  fetUseriInfoHandle() {
    fetchUserCenter().then(
      ({
        userInfo,
        countsData,
        orderTagInfos: orderInfo,
        customerServiceInfo,
      }) => {
        // eslint-disable-next-line no-unused-expressions
        menuData?.[0].forEach((v) => {
          countsData.forEach((counts) => {
            if (counts.type === v.type) {
              // eslint-disable-next-line no-param-reassign
              v.tit = counts.num;
            }
          });
        });
        const info = orderTagInfos.map((v, index) => ({
          ...v,
          ...orderInfo[index],
        }));
        this.setData({
          userInfo,
          menuData,
          orderTagInfos: info,
          customerServiceInfo,
          currAuthStep: 2,
        });
        wx.stopPullDownRefresh();
      },
    );
  },

  onClickCell({ currentTarget }) {
    const { type } = currentTarget.dataset;

    switch (type) {
      case 'address': {
        wx.navigateTo({ url: '/pages/usercenter/address/list/index' });
        break;
      }
      case 'service': {
        this.openMakePhone();
        break;
      }
      case 'help-center': {
        Toast({
          context: this,
          selector: '#t-toast',
          message: '你点击了帮助中心1',
          icon: '',
          duration: 1000,
        });
        break;
      }
      case 'point': {
        Toast({
          context: this,
          selector: '#t-toast',
          message: '你点击了积分菜单',
          icon: '',
          duration: 1000,
        });
        break;
      }
      case 'coupon': {
        wx.navigateTo({ url: '/pages/coupon/coupon-list/index' });
        break;
      }
      default: {
        Toast({
          context: this,
          selector: '#t-toast',
          message: '未知跳转',
          icon: '',
          duration: 1000,
        });
        break;
      }
    }
  },

  jumpNav(e) {
    const status = e.detail.tabType;

    if (status === 0) {
      wx.navigateTo({ url: '/pages/order/after-service-list/index' });
    } else {
      wx.navigateTo({ url: `/pages/order/order-list/index?status=${status}` });
    }
  },

  jumpAllOrder() {
    wx.navigateTo({ url: '/pages/order/order-list/index' });
  },

  openMakePhone() {
    this.setData({ showMakePhone: true });
  },

  closeMakePhone() {
    this.setData({ showMakePhone: false });
  },

  call() {
    wx.makePhoneCall({
      phoneNumber: this.data.customerServiceInfo.servicePhone,
    });
  },

  gotoUserEditPage() {
    const { currAuthStep } = this.data;
    console.log(currAuthStep);
    if (currAuthStep === 2) {
      // 取消跳转
      // wx.navigateTo({ url: '/pages/usercenter/person-info/index' });
    } else {

      // 注意使用函数的写法，避免出现错误
      let userProFile = this.getUserProfile();
      let loginCode = this.getLoginCode();
      loginCode
        .then(code => {
          return code;
        })
        .then(logCode => {
          return new Promise((resolve, reject) => {
            userProFile
              .then(res => {
                resolve({ logCode, iv: res.iv, rawData: res.rawData, encryptedData: res.encryptedData, signature: res.signature });
              })
              .catch(err => {
                reject(err);
              });
          });
        })
        .then(res => {
          // 入库
          var that = this;
          wx.request ({
            url: 'https://mung.lvdouiot.com/api/login', // api
            data: {
              encryptedData: res.encryptedData,
              iv: res.iv,
              code: res.logCode.code,  //用户登录凭证，有效期5分钟
              userInfo: res.rawData
            },
            header: {
              'content-type': 'application/json', //请求头
            },
            success: function (rres) {
                //json转化
                //const result = JSON.parse (rres);
                console.log("res:", rres.data)
                // 将用户授权信息存储到本地 还有openid
                wx.setStorageSync('userinfo', res.rawData)
                wx.setStorageSync('openid', rres.data.openid)
                // 将后端返回的token存储到本地
                //wx.setStorageSync('token', rres.data.token)
                that.setData({
                  userInfo: {
                    nickName: JSON.parse(res.rawData).nickName,
                    avatarUrl:JSON.parse(res.rawData).avatarUrl,
                  },
                  currAuthStep:2,
                })  
            },
            fail: function (rres) {
              console.log (rres);
            },
          })

        })
        .catch(err => {
          console.log('userProfile-err', err);
        });

      // this.fetUseriInfoHandle();

    }
  },

  getVersionInfo() {
    const versionInfo = wx.getAccountInfoSync();
    const { version, envVersion = __wxConfig } = versionInfo.miniProgram;
    this.setData({
      versionNo: envVersion === 'release' ? version : envVersion,
    });
  },

  getUserProfile() {
    return new Promise((resolve, reject) => {
      wx.getUserProfile({
        desc: '获取你的昵称、头像、地区及性别',
        success: userRes => {
          console.log('getUserProfile-res', userRes);
          resolve(userRes);
        },
        fail: userErr => {
          wx.showToast({
            title: '未授权！',
          });
          console.log('getUserProfile-err', userErr);
          reject(userErr);
        }
      });
    });
  },
  getLoginCode() {
    return new Promise((resolve, reject) => {
      wx.login({
        provider: 'weixin',
        success: res => {
          // if (res.code) {
          //   wx.request ({
          //     url: 'https://mung.lvdouiot.com/api/login', //测试api
          //     data: {
          //       code: res.code, //用户登录凭证，有效期5分钟
          //     },
          //     header: {
          //       'content-type': 'application/json', //请求头
          //     },
          //     success: function (rres) {
          //         //json转化
          //         //const result = JSON.parse (rres);
          //         //解构赋值
          //         const {openid, session_key, unionid, errorcode} = rres.data;
          //         console.log(rres);
          //         //数据缓存
          //         wx.setStorageSync('openid', openid)
          //         wx.setStorageSync('session_key', session_key)
          //         wx.setStorageSync('unionid', unionid)

          //     },
          //     fail: function (rres) {
          //       console.log (rres);
          //     },
          //   })
          // }
          console.log('login-->res', res);
          resolve(res);
        }
      });

    });
  },


});
