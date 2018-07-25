    var id = null;
    var msg = {
        alert: function(t) {
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
        fire: function(obj) {
            var evt = ['webkitAnimationEnd', 'mozAnimationEnd', 'msAnimationEnd', 'oAnimationEnd', 'animationEnd'];
            for (var i = 0; i < evt.length; i++) {
                obj.addEventListener(evt[i], function() {
                    this.parentNode.removeChild(this);
                });
            }
        }
    };
    $('.play').on('click', function() {
        var text = $('#sure').val();
        if (id != null) {
            if (text != '') {
                $.ajax({
                    url: '../../TestWeb/ajax/AIone.aspx?id=' + id + '&text=' + text,
                    type: 'GET',
                    success: function(res) {
                        console.log(res)
                        $('audio')[0].src = './images/1.mp3';
                        $('audio')[0].play();
                    },
                    error: function() {
                        msg.alert('网络连接不稳……');
                    }
                })
            } else {
                msg.alert('请输入内容')
            }
        } else {
            msg.alert('请选择性别')
        }

    })
    $('.add').on('click', function() {
        $('.giftbox').show(200);
    })
    $('.no').on('click', function() {
        $('.giftbox').hide(200);
    })
    $('.giftboy').on('click', function() {
        id = 2;
        $('#nan').val(id);
        $('#nv').val('null');
        $(this).find('i').css({
            'background': 'url("/static/images/man2.png")',
            "background-size": 'cover'
        });
        $(this).css('color', '#fa404d');
        $('.giftgirl').find('i').css({
            'background': 'url("/static/images/woman2.png")',
            "background-size": 'cover',
        });
        $('.giftgirl').css('color', '');
    })
    $('.giftgirl').on('click', function() {
        id = 4;
        $('#nan').val(id);
        $(this).find('i').css({
            'background': 'url("/static/images/woman1.png")',
            "background-size": 'cover',
            'color': '#fa404d'
        });
        $(this).css('color', '#fa404d');
        $('.giftboy').find('i').css({
            'background': 'url("/static/images/man1.png")',
            "background-size": 'cover',
        });
        $('.giftboy').css('color', '');
    })