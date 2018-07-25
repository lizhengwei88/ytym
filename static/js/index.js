var imgs = [];
$(document).ready()
{
    $.ajax({
        url: "/ajax",
        type: "Post",
        dataType: "Text",
        data: {
			       "type": "1"
		},
        success: function (data) {
			alert(data)
            if (data != "") {
			alert(data)
                var result = eval('(' + data + ')');
				alert(result)
                for (var i = 0; i < result.length; i++) {
					alert(result[i])
                    var name = result[i].cpname;
                    var pic = result[i].upic;
                    var img = { "name": name, "pic": pic };
                    imgs.push(img);
                }

                //z-index的值
                var z = 999999;
                //显示第几张图片
                var index = 0;
                var box = document.getElementById('box')

                boom(2, 2)
                //l 传进来几行；t传进来几列;
                var box = document.getElementById('box');
                function boom(l, t) {
                    $('#name').html(imgs[index].name);
                    $('#price').html(imgs[index].price);
                    $('#yprice').html(imgs[index].yprice);
                    $('#sail').html(imgs[index].sail);
                    if (imgs[index].like == 'like') {
                        $('#like').css({
                            'background': 'url("static/images/fa.png")',
                            'background-size': 'cover'
                        });
                    } else {
                        $('#like').css({
                            'background': 'url("static/images/aixin.png")',
                            'background-size': 'cover'
                        });
                    }
                    //创建一个新的div
			 
                    var oParentNode = document.createElement("div");
                    $(oParentNode).html(
                        `<div id="mark">
                                <img src="static/images/xin.png" alt="">
                            </div>`);
                    oParentNode.style.background = "url(" + imgs[index].pic + ") no-repeat";
                    oParentNode.style.backgroundSize = "89% 82%";
                    oParentNode.style.backgroundPosition = 'center';
                    //设置z-index的值
                    oParentNode.style.zIndex = z;
                    z--;
                    box.appendChild(oParentNode);
                    if (arguments[2] == 'left') {
                        $(oParentNode).css('marginLeft', '-6rem').animate({
                            'marginLeft': '0rem'
                        }, 400);
                    }
                    $("#box").on("touchend", function (e) {
                        // 判断默认行为是否可以被禁用
                        if (e.cancelable) {
                            // 判断默认行为是否已经被禁用
                            if (!e.defaultPrevented) {
                                e.preventDefault();
                            }
                        }
                        moveEndX = e.originalEvent.changedTouches[0].pageX,
                            moveEndY = e.originalEvent.changedTouches[0].pageY,
                            X = moveEndX - startX,
                            Y = moveEndY - startY;
                        //左滑
                        if (X > 0) {
                            index--;
                            if (index < 0) {
                                index = imgs.length - 1;
                            }
                            index == imgs.length && (index = 0);
                            $('#box').off('touchend');
                            boom(l, t, 'left');
                            $(oParentNode).animate({
                                'marginLeft': '0rem',
                                'opacity': 0
                            }, '1000');
                            setTimeout(function () {
                                oParentNode.remove();
                            }, 300)
                        }
                            //右滑
                        else if (X < 0) {
                            index++;
                            index == imgs.length && (index = 0);
                            $('#box').off('touchend');
                            boom(l, t);
                            $(oParentNode).animate({
                                'marginLeft': '-6rem',
                                'opacity': 0
                            }, '800');
                            setTimeout(function () {
                                oParentNode.remove();
                            }, 300)
                        }
                            //下滑
                        else if (Y > 0) {

                        }
                            //上滑
                        else if (Y < 0) {

                        }
                            //单击
                        else {
                            //alert('单击');
                        }
                        //$('#box').off('touchend')
                    });
                }
                $('#like').on('click', function () {
                    var that = this;
                    if (imgs[index].like == 'like') {
                        //取消喜欢
                        var customid = imgs[index].ID;
                        $.ajax({
                            url: "Index.aspx",
                            type: "Post",
                            dataType: "Text",
                            data: {
                                "type": "3", "customid": customid
                            },
                            success: function (data) {
                                if (data == "0") {
                                    msg.alert('取消喜欢失败！')
                                }
                                else {
                                    $(that).css({
                                        'background': 'url("static/images/aixin.png")',
                                        'background-size': 'cover'
                                    })
                                    msg.alert('取消喜欢！')
                                    imgs[index].like = 'unlike';
                                }
                            }
                        })
                    }
                    else {
                        //加入喜欢
                        var customid = imgs[index].ID;
                        $.ajax({
                            url: "Index.aspx",
                            type: "Post",
                            dataType: "Text",
                            data: {
                                "type": "2", "customid": customid
                            },
                            success: function (data) {
                                if (data == "0") {
                                    msg.alert('加入喜欢失败！')
                                }
                                else {
                                    $(that).css({
                                        'background': 'url("static/images/fa.png")',
                                        'background-size': 'cover'
                                    })
                                    msg.alert('加入喜欢！')
                                    imgs[index].like = 'like';
                                }
                            }
                        })
                    }
                });
                var time = null;
                $('#dislike').on('click', function () {
                    if (!time) {
                        index++;
                        index == imgs.length && (index = 0);
                        $('#box').off('touchend');
                        boom('l', 't');
                        $('#box>div').eq(0).animate({
                            'marginLeft': '-6rem',
                            'opacity': 0
                        }, '800');
                        setTimeout(function () {
                            $('#box>div').eq(0).remove();
                            time = null;
                        }, 300)
                    }

                })
            }
        }
    })
}

var msg = {
    alert: function (t) {
        var alt = document.querySelector('.alert');
        /*1：移除原有alert，2：不移除，直接替换文字。*/
        if (alt) { /*alt.parentNode.removeChild(alt);*/
            alt.innerHTML = t;
            return;
        };
        var h = document.createElement('div');
        h.innerHTML = t || '为什么不说话 (╰_╯)';
        h.setAttribute('class', 'alert');
        document.querySelector('body').appendChild(h);
        msg.fire(h);
    },
    fire: function (obj) {
        var evt = ['webkitAnimationEnd', 'mozAnimationEnd', 'msAnimationEnd', 'oAnimationEnd', 'animationEnd'];
        for (var i = 0; i < evt.length; i++) {
            obj.addEventListener(evt[i], function () {
                this.parentNode.removeChild(this);
            });
        }
    }
};
$("#box").on("touchstart", function (e) {
    // 判断默认行为是否可以被禁用
    if (e.cancelable) {
        // 判断默认行为是否已经被禁用
        if (!e.defaultPrevented) {
            e.preventDefault();
        }
    }
    startX = e.originalEvent.changedTouches[0].pageX,
        startY = e.originalEvent.changedTouches[0].pageY;
});