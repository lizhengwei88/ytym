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