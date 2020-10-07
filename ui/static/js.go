// Code generated by go generate; DO NOT EDIT.

package static // import "miniflux.app/ui/static"

var Javascripts = map[string]string{
	"app":            `!function(){'use strict';class a{static isVisible(a){return a.offsetParent!==null}static openNewTab(b){let a=window.open("");a.opener=null,a.location=b,a.focus()}static scrollPageTo(a){let d=window.pageYOffset,b=document.documentElement.clientHeight,c=d+b,e=a.offsetTop+a.offsetHeight;(c-e<0||c-a.offsetTop>b)&&window.scrollTo(0,a.offsetTop-10)}static getVisibleElements(c){let a=document.querySelectorAll(c),b=[];for(let c=0;c<a.length;c++)this.isVisible(a[c])&&b.push(a[c]);return b}static findParent(a,b){for(;a&&a!==document;a=a.parentNode)if(a.classList.contains(b))return a;return null}static hasPassiveEventListenerOption(){var b=!1,a;try{a=Object.defineProperty({},"passive",{get:function(){b=!0}}),window.addEventListener("test",a,a),window.removeEventListener("test",a,a)}catch(a){b=!1}return b}}class N{constructor(){this.reset()}reset(){this.touch={start:{x:-1,y:-1},move:{x:-1,y:-1},element:null}}calculateDistance(){if(this.touch.start.x>=-1&&this.touch.move.x>=-1){let a=Math.abs(this.touch.move.x-this.touch.start.x),b=Math.abs(this.touch.move.y-this.touch.start.y);if(a>30&&b<70)return this.touch.move.x-this.touch.start.x}return 0}findElement(b){return b.classList.contains("touch-item")?b:a.findParent(b,"touch-item")}onTouchStart(a){if(a.touches===void 0||a.touches.length!==1)return;this.reset(),this.touch.start.x=a.touches[0].clientX,this.touch.start.y=a.touches[0].clientY,this.touch.element=this.findElement(a.touches[0].target)}onTouchMove(a){if(a.touches===void 0||a.touches.length!==1||this.element===null)return;this.touch.move.x=a.touches[0].clientX,this.touch.move.y=a.touches[0].clientY;let b=this.calculateDistance(),c=Math.abs(b);if(c>0){let d=1-(c>75?.9:c/75*.9),e=b>75?75:b<-75?-75:b;this.touch.element.style.opacity=d,this.touch.element.style.transform="translateX("+e+"px)",a.preventDefault()}}onTouchEnd(a){if(a.touches===void 0)return;if(this.touch.element!==null){let a=Math.abs(this.calculateDistance());a>75&&n(this.touch.element),this.touch.element.style.opacity=1,this.touch.element.style.transform="none"}this.reset()}listen(){let e=document.querySelectorAll(".touch-item"),c=a.hasPassiveEventListenerOption();e.forEach(a=>{a.addEventListener("touchstart",a=>this.onTouchStart(a),!!c&&{passive:!0}),a.addEventListener("touchmove",a=>this.onTouchMove(a),!!c&&{passive:!1}),a.addEventListener("touchend",a=>this.onTouchEnd(a),!!c&&{passive:!0}),a.addEventListener("touchcancel",()=>this.reset(),!!c&&{passive:!0})});let d=document.querySelector(".entry-content");if(d){let a={previous:null,next:null};const e=(c,d)=>{const e=a[c];e===null?a[c]=setTimeout(()=>{a[c]=null},200):(d.preventDefault(),b(c))};d.addEventListener("touchend",a=>{a.changedTouches[0].clientX>=d.offsetWidth/2?e("next",a):e("previous",a)},!!c&&{passive:!1}),d.addEventListener("touchmove",b=>{Object.keys(a).forEach(b=>a[b]=null)})}}}class M{constructor(){this.queue=[],this.shortcuts={},this.triggers=[]}on(a,b){this.shortcuts[a]=b,this.triggers.push(a.split(" ")[0])}listen(){document.onkeydown=a=>{let b=this.getKey(a);if(this.isEventIgnored(a,b)||this.isModifierKeyDown(a))return;a.preventDefault(),this.queue.push(b);for(let c in this.shortcuts){let d=c.split(" ");if(d.every((a,b)=>a===this.queue[b])){this.queue=[],this.shortcuts[c](a);return}if(d.length===1&&b===d[0]){this.queue=[],this.shortcuts[c](a);return}}this.queue.length>=2&&(this.queue=[])}}isEventIgnored(a,b){return a.target.tagName==="INPUT"||a.target.tagName==="TEXTAREA"||this.queue.length<1&&!this.triggers.includes(b)}isModifierKeyDown(a){return a.getModifierState("Control")||a.getModifierState("Alt")||a.getModifierState("Meta")}getKey(b){const a={Esc:'Escape',Up:'ArrowUp',Down:'ArrowDown',Left:'ArrowLeft',Right:'ArrowRight'};for(let c in a)if(a.hasOwnProperty(c)&&c===b.key)return a[c];return b.key}}class d{constructor(a){this.callback=null,this.url=a,this.options={method:"POST",cache:"no-cache",credentials:"include",body:null,headers:new Headers({"Content-Type":"application/json","X-Csrf-Token":this.getCsrfToken()})}}withHttpMethod(a){return this.options.method=a,this}withBody(a){return this.options.body=JSON.stringify(a),this}withCallback(a){return this.callback=a,this}getCsrfToken(){let a=document.querySelector("meta[name=X-CSRF-Token]");return a!==null?a.getAttribute("value"):""}execute(){fetch(new Request(this.url,this.options)).then(a=>{this.callback&&this.callback(a)})}}class g{static exists(){return document.getElementById("modal-container")!==null}static open(c){if(g.exists())return;let a=document.createElement("div");a.id="modal-container",a.appendChild(document.importNode(c,!0)),document.body.appendChild(a);let b=document.querySelector("a.btn-close-modal");b!==null&&(b.onclick=a=>{a.preventDefault(),g.close()})}static close(){let a=document.getElementById("modal-container");a!==null&&a.parentNode.removeChild(a)}}function c(a,b,c){let d=document.querySelectorAll(a);d.forEach(a=>{a.onclick=a=>{c||a.preventDefault(),b(a)}})}function L(){let b=document.querySelector(".header nav ul");a.isVisible(b)?b.style.display="none":b.style.display="block";let c=document.querySelector(".header .search");a.isVisible(c)?c.style.display="none":c.style.display="block"}function K(b){let a=b.target;a.tagName==="A"?window.location.href=a.getAttribute("href"):window.location.href=a.querySelector("a").getAttribute("href")}function J(){let a=document.querySelectorAll("form");a.forEach(a=>{a.onsubmit=()=>{let b=a.querySelector("button");b&&(b.innerHTML=b.dataset.labelLoading,b.disabled=!0)}})}function q(b){b.preventDefault(),b.stopPropagation();let c=document.querySelector(".search-toggle-switch");c&&(c.style.display="none");let d=document.querySelector(".search-form");d&&(d.style.display="block");let a=document.getElementById("search-input");a&&(a.focus(),a.value="")}function G(){let a=document.getElementById("keyboard-shortcuts");a!==null&&g.open(a.content)}function p(){let d=a.getVisibleElements(".items .item"),c=[];d.forEach(a=>{a.classList.add("item-status-read"),c.push(parseInt(a.dataset.id,10))}),c.length>0&&j(c,"read",()=>{let a=document.querySelector("a[data-action=markPageAsRead]"),c=!1;a&&(c=a.dataset.showOnlyUnread||!1),c?window.location.reload():b("next",!0)})}function m(b){let c=!b,a=i(b);a&&(n(a,c),e()&&a.classList.contains('current-item')&&k())}function n(b,d){let g=parseInt(b.dataset.id,10),a=b.querySelector("a[data-toggle-status]"),c=a.dataset.value,e=c==="read"?"unread":"read";j([g],e),c==="read"?(a.innerHTML='<span class="icon-label">'+a.dataset.labelRead+'</span>',a.dataset.value="unread",d&&f(a.dataset.toastUnread)):(a.innerHTML='<span class="icon-label">'+a.dataset.labelUnread+'</span>',a.dataset.value="read",d&&f(a.dataset.toastRead)),b.classList.contains("item-status-"+c)&&(b.classList.remove("item-status-"+c),b.classList.add("item-status-"+e))}function E(a){if(a.classList.contains("item-status-unread")){a.classList.remove("item-status-unread"),a.classList.add("item-status-read");let b=parseInt(a.dataset.id,10);j([b],"read")}}function C(){let b=document.body.dataset.refreshAllFeedsUrl,a=new d(b);a.withCallback(()=>{window.location.reload()}),a.withHttpMethod("GET"),a.execute()}function j(c,b,e){let f=document.body.dataset.entriesStatusUrl,a=new d(f);a.withBody({entry_ids:c,status:b}),a.withCallback(e),a.execute(),b==="read"?H(1):I(1)}function r(a){let c=!a,b=i(a);b&&B(b.querySelector("a[data-save-entry]"),c)}function B(a,c){if(!a)return;if(a.dataset.completed)return;let e=a.innerHTML;a.innerHTML='<span class="icon-label">'+a.dataset.labelLoading+'</span>';let b=new d(a.dataset.saveUrl);b.withCallback(()=>{a.innerHTML=e,a.dataset.completed=!0,c&&f(a.dataset.toastDone)}),b.execute()}function t(a){let c=!a,b=i(a);b&&O(b,c)}function O(e,b){let a=e.querySelector("a[data-toggle-bookmark]");if(!a)return;a.innerHTML='<span class="icon-label">'+a.dataset.labelLoading+'</span>';let c=new d(a.dataset.bookmarkUrl);c.withCallback(()=>{a.dataset.value==="star"?(a.innerHTML='<span class="icon-label">'+a.dataset.labelStar+'</span>',a.dataset.value="unstar",b&&f(a.dataset.toastUnstar)):(a.innerHTML='<span class="icon-label">'+a.dataset.labelUnstar+'</span>',a.dataset.value="star",b&&f(a.dataset.toastStar))}),c.execute()}function v(){if(e())return;let a=document.querySelector("a[data-fetch-content-entry]");if(!a)return;let c=a.innerHTML;a.innerHTML='<span class="icon-label">'+a.dataset.labelLoading+'</span>';let b=new d(a.dataset.fetchContentUrl);b.withCallback(b=>{a.innerHTML=c,b.json().then(a=>{a.hasOwnProperty("content")&&(document.querySelector(".entry-content").innerHTML=a.content)})}),b.execute()}function w(d){let b=document.querySelector(".entry h1 a");if(b!==null){d?window.location.href=b.getAttribute("href"):a.openNewTab(b.getAttribute("href"));return}let c=document.querySelector(".current-item a[data-original-link]");if(c!==null){a.openNewTab(c.getAttribute("href"));let b=document.querySelector(".current-item");document.location.href!=document.querySelector('a[data-page=starred]').href&&k(),E(b)}}function x(b){if(e()){let b=document.querySelector(".current-item a[data-comments-link]");b!==null&&a.openNewTab(b.getAttribute("href"))}else{let c=document.querySelector("a[data-comments-link]");if(c!==null){b?window.location.href=c.getAttribute("href"):a.openNewTab(c.getAttribute("href"));return}}}function A(){let a=document.querySelector(".current-item .item-title a");a!==null&&(window.location.href=a.getAttribute("href"))}function z(){let a=document.querySelectorAll("[data-action=remove-feed]");if(a.length===1){let b=a[0],c=new d(b.dataset.url);c.withCallback(()=>{b.dataset.redirectUrl?window.location.href=b.dataset.redirectUrl:window.location.reload()}),c.execute()}}function b(b,c){let a=document.querySelector("a[data-page="+b+"]");a?document.location.href=a.href:c&&window.location.reload()}function h(){e()?F():b("previous")}function l(){e()?k():b("next")}function D(){s()?o():b('feeds')}function o(){if(s()){let a=document.querySelector("span.entry-website a");a!==null&&(window.location.href=a.href)}else{let a=document.querySelector(".current-item a[data-feed-link]");a!==null&&(window.location.href=a.getAttribute("href"))}}function F(){let b=a.getVisibleElements(".items .item");if(b.length===0)return;if(document.querySelector(".current-item")===null){b[0].classList.add("current-item"),b[0].querySelector('.item-header a').focus();return}for(let c=0;c<b.length;c++)if(b[c].classList.contains("current-item")){b[c].classList.remove("current-item");let d;c-1>=0?d=b[c-1]:d=b[b.length-1],d.classList.add("current-item"),a.scrollPageTo(d),d.querySelector('.item-header a').focus();break}}function k(){let b=a.getVisibleElements(".items .item");if(b.length===0)return;if(document.querySelector(".current-item")===null){b[0].classList.add("current-item"),b[0].querySelector('.item-header a').focus();return}for(let c=0;c<b.length;c++)if(b[c].classList.contains("current-item")){b[c].classList.remove("current-item");let d;c+1<b.length?d=b[c+1]:d=b[0],d.classList.add("current-item"),a.scrollPageTo(d),d.querySelector('.item-header a').focus();break}}function H(a){y(b=>b-a)}function I(a){y(b=>b+a)}function y(a){let b=document.querySelectorAll("span.unread-counter");if(b.forEach(b=>{let c=parseInt(b.textContent,10);b.innerHTML=a(c)}),window.location.href.endsWith('/unread')){let b=parseInt(document.title.split('(')[1],10),c=a(b);document.title=document.title.replace(/(.*?)\(\d+\)(.*?)/,function(d,a,b,e,f){return a+'('+c+')'+b})}}function s(){return document.querySelector("section.entry")!==null}function e(){return document.querySelector(".items")!==null}function i(b){return e()?b?a.findParent(b,"item"):document.querySelector(".current-item"):document.querySelector(".entry")}function u(a,f){a.tagName!='A'&&(a=a.parentNode),a.style.display="none";let e=a.parentNode,b=document.createElement("span"),c=document.createElement("a");c.href="#",c.appendChild(document.createTextNode(a.dataset.labelYes)),c.onclick=d=>{d.preventDefault();let c=document.createElement("span");c.className="loading",c.appendChild(document.createTextNode(a.dataset.labelLoading)),b.remove(),e.appendChild(c),f(a.dataset.url,a.dataset.redirectUrl)};let d=document.createElement("a");d.href="#",d.appendChild(document.createTextNode(a.dataset.labelNo)),d.onclick=c=>{c.preventDefault(),a.style.display="inline",b.remove()},b.className="confirm",b.appendChild(document.createTextNode(a.dataset.labelQuestion+" ")),b.appendChild(c),b.appendChild(document.createTextNode(", ")),b.appendChild(d),e.appendChild(b)}function f(a){if(!a)return;document.querySelector('.toast-wrap .toast-msg').innerHTML=a;let b=document.querySelector('.toast-wrap');b.classList.remove('toastAnimate'),setTimeout(function(){b.classList.add('toastAnimate')},100)}document.addEventListener("DOMContentLoaded",function(){if(J(),!document.querySelector("body[data-disable-keyboard-shortcuts=true]")){let a=new M;a.on("g u",()=>b("unread")),a.on("g b",()=>b("starred")),a.on("g h",()=>b("history")),a.on("g f",()=>D()),a.on("g c",()=>b("categories")),a.on("g s",()=>b("settings")),a.on("ArrowLeft",()=>h()),a.on("ArrowRight",()=>l()),a.on("k",()=>h()),a.on("p",()=>h()),a.on("j",()=>l()),a.on("n",()=>l()),a.on("h",()=>b("previous")),a.on("l",()=>b("next")),a.on("o",()=>A()),a.on("v",()=>w()),a.on("V",()=>w(!0)),a.on("c",()=>x()),a.on("C",()=>x(!0)),a.on("m",()=>m()),a.on("A",()=>p()),a.on("s",()=>r()),a.on("d",()=>v()),a.on("f",()=>t()),a.on("F",()=>o()),a.on("R",()=>C()),a.on("?",()=>G()),a.on("#",()=>z()),a.on("/",a=>q(a)),a.on("Escape",()=>g.close()),a.listen()}let a=new N;if(a.listen(),c("a[data-save-entry]",a=>r(a.target)),c("a[data-toggle-bookmark]",a=>t(a.target)),c("a[data-fetch-content-entry]",()=>v()),c("a[data-action=search]",a=>q(a)),c("a[data-action=markPageAsRead]",()=>u(event.target,()=>p())),c("a[data-toggle-status]",a=>m(a.target)),c("a[data-confirm]",a=>u(a.target,(c,a)=>{let b=new d(c);b.withCallback(()=>{a?window.location.href=a:window.location.reload()}),b.execute()})),document.documentElement.clientWidth<600&&(c(".logo",()=>L()),c(".header nav li",a=>K(a))),"serviceWorker"in navigator){let a=document.getElementById("service-worker-script");a&&navigator.serviceWorker.register(a.src)}window.addEventListener('beforeinstallprompt',c=>{c.preventDefault();let a=c;const b=document.getElementById('prompt-home-screen');if(b){b.style.display="block";const c=document.getElementById('btn-add-to-home-screen');c&&c.addEventListener('click',c=>{c.preventDefault(),a.prompt(),a.userChoice.then(()=>{a=null,b.style.display="none"})})}})})}()`,
	"service-worker": `self.addEventListener("fetch",a=>{a.request.url.includes("/feed/icon/")&&a.respondWith(caches.open("feed_icons").then(b=>b.match(a.request).then(c=>c||fetch(a.request).then(c=>(b.put(a.request,c.clone()),c)))))})`,
}

var JavascriptsChecksums = map[string]string{
	"app":            "4fbd2da66463adba7faa2179665e33e6146a9080d799465453d9d408734c257c",
	"service-worker": "730f10dc6a52e0bd9271da0c3b0103368893f3feb0a092fd585ac5b7abedb4ac",
}
