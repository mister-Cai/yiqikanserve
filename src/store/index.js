import Vue from 'vue'
import Vuex from 'vuex'
import { app } from "../main.js"
Vue.use(Vuex)

import detectEthereumProvider from "@metamask/detect-provider";
import Web3 from 'web3';
import { SaleContractAbi, saleContractAddr, USDTContractAbi, usdtContractAddr } from "@/utils/contractAbi"
let BscChainId = 97
export default new Vuex.Store({
	state: {
		//这里放全局参数
		saleContractAddr,
		usdtContract: '',
		saleContract: '',
		web3: '',
		loading: false,
		loadtext: "",
		ismetamask: true,
		wellatmsg: '',
		chainmsg: "",
		myaccount: '',
		approvedAmount: 0,
		balance: 0,
		langue: '中文',
		current:0,
		superList5:[],
		timelist5:[],
		superList:[],
		timelist:[],
		timeamount:0,
		inttimers:'',
		inttimers1:'',
		timers:[0,0,0],
		timers1:[0,0,0]
	},
	mutations: {
		setloading(state, flag) {
			state.loading = flag
		},
		setloadingtext(state, text) {
			state.loadtext = text
		},
		//这里是set方法
		handleChainChange(state, chainId) {
			console.log("current chainid=", chainId);
			if (chainId != BscChainId) {
				console.log("Chain is not supported.");
				state.ismetamask = false
				if (state.chainmsg) { state.chainmsg.close() }
				state.chainmsg = app.$message({
					duration: 0,
					showClose: true,
					type: 'error',
					message: "该网络不支持"
				});
				// 提示用户选择正确的网网络并禁用相关按钮
			} else {
				state.ismetamask = true
				this._actions.getdata[0]();
				// this._actions.usdtBalance[0]()
				if (state.chainmsg)
					state.chainmsg.close()
			}
		},
		handleAccountsChanged(state, accounts) {
			if (accounts.length === 0) {
				// MetaMask is locked or the user has not connected any accounts
				state.ismetamask = false
				state.wellatmsg = app.$message({
					duration: 0,
					showClose: true,
					type: 'error',
					message: "请安装Metamask"
				});
				console.log('Please connect to MetaMask.');
				// 窗提示用户安装metamask
			} else if (accounts[0] !== state.myaccount) {
				if (state.wellatmsg) { state.wellatmsg.close() }
				console.log('check account success');
				state.myaccount = accounts[0]
				if (ethereum.networkVersion != BscChainId) {
					if (state.chainmsg) { state.chainmsg.close() }
					state.ismetamask = false
					state.chainmsg = app.$message({
						duration: 0,
						showClose: true,
						type: 'error',
						message: "该网络不支持"
					});
				} else {
					state.ismetamask = true
					// 检查当前用户是否已授权
					this._actions.getdata[0]();
					//   this._actions.usdtBalance[0]()
					if (state.chainmsg)
						state.chainmsg.close();
				}
				// 显示当前连接的钱包帐户
			}
		},
		gettime(state){
			state.saleContract.methods.SUPER_LOTTERY_DURATION().call({ from: state.myaccount }).then(res=>{
				console.log('SUPER_LOTTERY_DURATION:'+res)
				state.saleContract.methods.getLastHeartBeat().call({ from: state.myaccount }).then(res1=>{
					console.log('getLastHeartBeat:'+res1)
					state.inttimers1=res-(new Date().getTime()/1000).toFixed()+res1*1
					if(state.inttimers1<=0){
						state.timers1=[0,0,0]
						return;
					}
					state.timers1[0]=Math.floor(state.inttimers1/3600)
					state.timers1[0]=Math.floor(state.inttimers1%3600/60)
					state.timers1[0]=(state.inttimers1%60)
					state.inttimers1=setInterval(()=>{
						if(state.timers1[2]*1-1==0){
							state.timers1[2]='60'
							state.timers1[2]='59'
						}else{
							state.timers1.splice(2,1,state.timers1[2]*1<11?'0'+(state.timers1[2]*1-1):state.timers1[2]*1-1+'') 
							return ;
						}
						if(state.timers1[1]==0){
							state.timers1[1]='60'
							state.timers1[1]='59'
						}else{
							state.timers1.splice(1,1,state.timers1[1]*1<11?'0'+(state.timers1[1]*1-1):state.timers1[1]*1-1+'')
							return ;
						}
						if(state.timers1[0]==0){
							state.timers1=[0,0,0]
							clearInterval(state.inttimers1)
						}else{
							state.timers1.splice(0,1,state.timers1[0]*1<11?'0'+(state.timers1[0]*1-1):state.timers1[0]*1-1+'')
							return ;
						}
						
					},1000)
				})
			})
			state.saleContract.methods.LAUNCH_TIME().call({ from: state.myaccount }).then(res=>{
				console.log('LAUNCH_TIME:'+res)
				state.saleContract.methods.LOTTERY_DURATION().call({ from: state.myaccount }).then(restime=>{
					console.log('LOTTERY_DURATION:'+restime)
					state.inttimers= Math.floor(restime-((new Date().getTime()/1000).toFixed()-res*1)%restime)
					if(state.inttimers==0){
						return;
					}
					state.timers[0]=Math.floor(state.inttimers/3600)
					state.timers[1]=Math.floor(state.inttimers%3600/60)
					state.timers[2]=(state.inttimers%60)
					state.inttimers=setInterval(()=>{
						if(state.timers[2]*1-1==0){
							state.timers[2]='60'
							state.timers[2]='59'
						}else{
							state.timers.splice(2,1,state.timers[2]*1<11?'0'+(state.timers[2]*1-1):state.timers[2]*1-1+'') 
							return ;
						}
						if(state.timers[1]==0){
							state.timers[1]='60'
							state.timers[1]='59'
						}else{
							state.timers.splice(1,1,state.timers[1]*1<11?'0'+(state.timers[1]*1-1):state.timers[1]*1-1+'')
							return ;
						}
						if(state.timers[0]==0){
							state.timers=[0,0,0]
							clearInterval(state.inttimers)
						}else{
							state.timers.splice(0,1,state.timers[0]*1<11?'0'+(state.timers[0]*1-1):state.timers[0]*1-1+'')
							return ;
						}
						
					},1000)
				})
			})
		}
	},
	getters: {
		// 这里是get方法  
	},
	actions: {
		async init({ commit, state }) {
			let provider = await detectEthereumProvider()
			if (provider) {
				state.web3 = new Web3(provider);
				state.saleContract = new state.web3.eth.Contract(SaleContractAbi, saleContractAddr)
				state.usdtContract = new state.web3.eth.Contract(USDTContractAbi, usdtContractAddr)
				state.loadtext = "连接中"
				state.loading = true
				await ethereum.request({ method: "eth_requestAccounts" }).then((res) => {
					console.log('getaccount success');
					commit('handleAccountsChanged', res)
					state.loading = false
				})
					.catch((err) => {
						// Some unexpected error.
						// For backwards compatibility reasons, if no accounts are available,
						// eth_accounts will return an empty array.
						console.error(err);
					});
				ethereum.on('accountsChanged', (res) => { commit('handleAccountsChanged', res) });
				ethereum.on('chainChanged', (chainId) => {
					commit('handleChainChange', state.web3.utils.hexToNumber(chainId));
				});
			} else {
				console.log("Please install MetaMask!");
				state.ismetamask = false
				state.wellatmsg = app.$message({
					duration: 0,
					showClose: true,
					type: 'error',
					message: "请安装Metamask"
				});
				// 弹窗提示用户安装Metamask
			}
		},
		getdata({ commit, state }){
			console.log(state.saleContract)
			state.usdtContract.methods
			.balanceOf("0x024b6b81ed8781a22c8153554d99f69f12f84459")
			.call({ from: "0x024b6b81ed8781a22c8153554d99f69f12f84459" })
			.then((result)=> {
			  state.balance=state.web3.utils.fromWei(result, 'ether')
			  console.log("current address, usdt balance =", state.web3.utils.fromWei(result, 'ether'), ' USDT');
			})
			.catch((err) => {
			  // If the request fails, the Promise will reject with an error.
			  console.error("checkApprove err, ", err);
			});
			state.saleContract.methods
			.getLotteryRound()
			.call({ from: state.myaccount })
			.then((result)=> {
				// state.current=result
				state.current=2
				console.log("getLotteryRound:",result)
				state.saleContract.methods //查看最新抢购地址
				.getLastBuyers(state.current)
				.call({ from: state.myaccount })
				.then((result)=> {
					console.log(1,result)
					for(let i in result){
						state.superList5.push({})
						state.saleContract.methods
						.buyin(state.current,result[i]).call({ from: state.myaccount }).then(res=>{
							state.superList5.splice(i,i+1,{add:result[i],mount:res})
						})
					}
				})
				.catch((err) => {
					// If the request fails, the Promise will reject with an error.
					console.error("checkApprove err, ", err);
				});
				state.saleContract.methods //查看交易,累计购买前五
				.getTopBuyers(result)
				.call({ from: state.myaccount })
				.then((result)=> {
					// 当前地址USDT余额
					console.log("getTopBuyers:",result)
					for(let i in result){
						state.timelist5.push({})
						state.saleContract.methods
						.buyin(state.current,result[i]).call({ from: state.myaccount }).then(res=>{
							state.timelist5.splice(i,i+1,{add:result[i],mount:res})
						})
					}
				})
				.catch((err) => {
					// If the request fails, the Promise will reject with an error.
					console.error("checkApprove err, ", err);
				});
				state.saleContract.getPastEvents('LotteryRewards',(error, events)=>{
					console.log('getPastEvents:',events)
					for(let i in events){
						state.timelist.push({})
						state.saleContract.methods
						.buyin(state.current,events[i]).call({ from: state.myaccount }).then(res=>{
							state.timeamount+=res
							state.timelist.splice(i,i+1,{add:events[i],mount:res})
						})
					}
				})
				state.saleContract.getPastEvents('SuperLotteryRewards',(error, events)=>{
					console.log('getPastEvents:',events)
					// state.superList=events
					let arr=[45,25,15,10,5] 
					for(let i in events){
						state.superList.push({add:events[i],mount:arr[i]})
					}
				})
			})
			.catch((err) => {
			  // If the request fails, the Promise will reject with an error.
			  console.error("checkApprove err, ", err);
			});
			commit('gettime')
		}
	},
	modules: {
		//这里是我自己理解的是为了给全局变量分组，所以需要写提前声明其他store文件，然后引入这里

	}

})
