// font-size blind link
$('.x2').on('click tap',function(){
	$('.blind-version').removeClass('x4font x8font');
	$('.blind-version').addClass('x2font');
	
	$.ajax({
		url: '/default/index/setfont',
		data: {
			font : 'x2font',
		},
		type: 'POST'
	});
	return false;
});
$('.x4').on('click tap',function(){
	$('.blind-version').removeClass('x2font x8font');
	$('.blind-version').addClass('x4font');
	
	$.ajax({
		url: '/default/index/setfont',
		data: {
			font : 'x4font',
		},
		type: 'POST'
	});

	return false;
});
$('.x8').on('click tap',function(){
	$('.blind-version').removeClass('x2font x4font');
	$('.blind-version').addClass('x8font');
	
	$.ajax({
		url: '/default/index/setfont',
		data: {
			font : 'x8font',
		},
		type: 'POST'
	});
	return false;
});

// special version
$('.specVersion').not('.temporary-text').on('click tap', function(){
	$('body').toggleClass('blind-version');
	
	$.ajax({
		url: '/default/index/setspecialview',		
		type: 'POST'
	});	
	return false;
});


// sticky menu
var Main_menu = {
	
	delay: 300,

	timer: {
		open: 0,
		hide: 0,
	},

	show_menu: function(elem){
		if($(elem).hasClass('submenu_li')){
			Main_menu.timer.open = setTimeout(function() { 
				var dataId = $($(elem).data('id'));
				$('.submenu_li,.submenu-left,.submenu-right,.nav-link__item').removeClass('active');
				dataId.addClass('active');
				$(elem).find('.nav-link__item').addClass('active');
			}, Main_menu.delay);

		}else{
			console.log(2);
			clearTimeout(Main_menu.timer.open);
			clearTimeout(Main_menu.timer.hide);
			var dataId = $('.submenu_li[data-id="#' + $(elem).attr('id') + '"]');
			dataId.find('.nav-link__item').addClass('active');
			$(elem).addClass('active');
		}
	},

	hide_menu: function(elem){
		if($(elem).hasClass('submenu_li')){
			var dataId = $($(elem).data('id'));
					//dataId.removeClass('active');
					//$(this).find('.nav-link__item').removeClass('active');	
				}else{
					var dataId = $('.submenu_li[data-id="#' + $(elem).attr('id') + '"]');
					dataId.find('.nav-link__item').removeClass('active');
					$(elem).removeClass('active');	
				}
			},

			init: function(){

				$( '.nav-left__menu,.nav-right__menu' ).mouseleave(function(event) {
					if($(event.toElement).parents('.logo-wrap').length || $(event.toElement).hasClass('logo-wrap'))
						return;
					$('.submenu_li,.submenu-left,.submenu-right,.nav-link__item').removeClass('active');

				});

				$('header').mouseleave(function(){
					clearTimeout(Main_menu.timer.open);
					clearTimeout(Main_menu.timer.hide);
					$('.submenu_li,.submenu-left,.submenu-right,.nav-link__item').removeClass('active');
				});

				$('.submenu_li, .submenu-left, .submenu-right').hover(
					function () {
						Main_menu.show_menu(this);
					}, 
					function () {
						//console.log(1);
						var elem = this;
						Main_menu.timer.hide = setTimeout(function() { 
							Main_menu.hide_menu(elem);
						}, Main_menu.delay);

					}
					);
			}
		}

		$(function(){
			$(window).scroll(function() {
				if($(this).scrollTop() >= 20 && $(window).width() > 768) {
					$('.logo-first').addClass('logo-scroll');
					$('.logo-second').removeClass('logo-scroll').addClass('logo-active');
					$('header').addClass('stickytop');
				}
				else{
					$('header').removeClass('stickytop');
					$('.logo-first').removeClass('logo-scroll').addClass('logo-active');
					$('.logo-second').removeClass('logo-active').addClass('logo-scroll');

				}
			});
			$(window).scroll();
		});

	
	$(document).ready(function(){
	


	// Search
	var search = $('.search-wrap');

	$('.search').on('click tap',function(){
		search.addClass('active');
		$('.search-wrap__item').find('input[type="text"]').focus();		
		return false;
	});
	$('.close').on('click tap',function(){
		search.removeClass('active');
		return false;
	});

	// gamburger-menu
	var gm = $('.gm-button');

	gm.on('click tap', function(){
		$(this).toggleClass('on');
		$('.mbl-container').toggleClass('active');
	});

	// Submenu
	Main_menu.init();

	// Pages 
	$('#tab-nav li').on('click tap', function () {
		var navPage = $('#tab-nav li'),
		content = $('.page-wrap'),
		dataId = $($(this).data('id'));

		navPage.removeClass('active').removeClass('active-nav');
		content.removeClass('active-block');

		$(this).addClass('active');
		dataId.addClass('active-block');
		
		$('html, body').animate({scrollTop: $('.page-nav').offset().top-70}, 250);
	});
	
	if($('.page-wrap').length && window.location.hash!==''){
		$('#tab-nav li[data-id="'+window.location.hash+'"]').trigger('click');
	}

	// spoiler
	var spoil = $('.spoiler-body');
	spoil.hide();
	$('.spoiler-title').click(function(e){
		if($(this).next().hasClass('spoiler-body'))
		{
			$(this).next().slideToggle(300);
			$(this).toggleClass('active');		
			return false;
		}
	});

	var spPage = $('.spoiler-body__page');
	spPage.hide();

	$('.spoiler-title__page').on('click',function(){
		$(this).next().slideToggle(300);
		$(this).toggleClass('active');
		return false;
	});


	// $('#accardeon .acc-head[href^="#"]').click(function(){
	// 	var target = $(this).attr('href');
	// 	setTimeout(function(){
	// 		$('html, body').animate({scrollTop: $(target).offset().top}, 250);
	// 		return false;
	// 	},550)

	// });

// pages spoiler
var mblContent = $('.mbl-content, .mbl-context ');

$('.mbl-wrapper__nav').each(function(i){
	if( $(this).hasClass('active')){			
		$(this).next().show();
	}
	else{
		mblContent.hide();
	}
});





$('.mbl-nav').on('click', function(){
	$(this).next().slideToggle(300);
	$(this).toggleClass('active');
});

$('.mbl-content__title').on('click', function(){
	$(this).next().slideToggle(300);
	$(this).toggleClass('active');

});


if($('#add-feedback-el-form, #subsrcibe-form').length )
{	
	$('body').on('keyup change focusout', '[data-valid]', function () {
		Validator.validate($(this));
	});

	$('[data-mask]').each(function(elem){
		$(this).inputmask($(this).attr('data-mask'));
	});

	$('body').on('click', '.feedbackbtn', function(){
		$(this).parents('form').submit();
	});

	$('#add-feedback-el-form').submit(function(){
		$('[data-valid]').each(function(index, element) {
			Validator.validate($(element));
		});

		if($('[data-valid].error-input').length > 0){
			$('html, body').animate({
				scrollTop: $(".error-input").eq(0).offset().top - 120
			}, 500);
			return false;
		}
	});
}
});

if($('#subsrcibe-form').length){
	$( "body" ).on('click', '#subsrcibe-form .subscribe-box__btn', function(){
		$('#subsrcibe-form').submit();
		return false;
	});

	$( "body" ).on('submit', '#subsrcibe-form', function(e){
		e.preventDefault();
		var $this=$(this),
		pattern = /^([a-z0-9_\.-])+@[a-z0-9-]+\.([a-z]{2,4}\.)?[a-z]{2,4}$/i;

		if(!$('#subscribe-email').val() || !pattern.test($('#subscribe-email').val())){
			$('#subscribe-email').focus();
			return false;
		}

		$.ajax({
			url: $this.attr('action'),
			data: $this.serialize(),
			type: 'POST',
			dataType: 'json',
			beforeSend: function(){
				$this.css('opacity', '0.6');
			},
			success: function(data){
				if(data.status=='0'){
					$this.find('.subscribe-box__item').html(data.msg).hide().fadeIn();
					$this.find('.subscribe-box__btn').remove();
				}else{
					$this.find('.subscribe-box__item').html('РџРѕРјРёР»РєР°').hide().fadeIn();
					$this.find('.subscribe-box__btn').remove();					
				}
				$this.css('opacity', '1');
			}
		});

	});
}

//===================SCROLL TOP ARROW==========================//
var ScrollTopArrow = (function(){
	var arrow = '<div class="button__scrollTop-arrow hidden"><span class="scrollTop-arrow__inner"></span></div>';
	var initialized = false;
	var init = function() {
		$('body').append(arrow);
		
		setTimeout(function(){
			var elem = $('.button__scrollTop-arrow');
				
			$(window).scroll(function() {
				var offset = window.scrollY;
				if(offset > 500) {
					elem.removeClass('hidden');
				} else {
					elem.addClass('hidden');					
				}
			});
			
			elem.on('click', function(){
				var self = this;
				$('html, body').animate({
					scrollTop: 0
				}, 500);
			});
		}, 0);
	}
	return {
		init: function() {
			if(!initialized) {
				initialized = true;
				return init();				
			}
			return;
		}
	}
}());

ScrollTopArrow.init();
//===================/.END SCROLL TOP ARROW==========================//

/**
|--------------------------------------------------
| SEARCH FORM MIDDLEWARE
|--------------------------------------------------
*/
$(document).ready(function() {
	var search_form_middleware = (function(){
		var search_form_header = $('.search-wrap__item form'), 
			search_form_page = $('.search-input-page form');

		search_form_header.add(search_form_page).on('submit', function(event) {
			event.preventDefault();

			var form_value = event.target[0].value.replace(/\//gi, '');
			
			window.location ='/search/list/q/' + form_value;
		});
	}());
});

/**
|--------------------------------------------------
| /.END SEARCH FORM MIDDLEWARE
|--------------------------------------------------
*/

$(document).ready(function() {
	$( ".shadowTop1, .shadowLine1, .shadowBottom1" ).hover(
	  function() {
		$( '.shadowTop1' ).removeClass("disabled");
		$( '.shadowLine1' ).removeClass("disabled");
		$( '.shadowBottom1' ).removeClass("disabled");
		
		$( '.shadowTop' ).addClass("disabled");
		$( '.shadowLine' ).addClass("disabled");
		$( '.shadowBottom' ).addClass("disabled");
	  }, function() {
		$( '.shadowTop1' ).addClass("disabled");
		$( '.shadowLine1' ).addClass("disabled");
		$( '.shadowBottom1' ).addClass("disabled");
		
		$( '.shadowTop' ).removeClass("disabled");
		$( '.shadowLine' ).removeClass("disabled");
		$( '.shadowBottom' ).removeClass("disabled");
	  }
	);
	
	$( ".shadowTop2, .shadowLine2, .shadowBottom2" ).hover(
	  function() {
		$( '.shadowTop2' ).removeClass("disabled");
		$( '.shadowLine2' ).removeClass("disabled");
		$( '.shadowBottom2' ).removeClass("disabled");
		
		$( '.shadowTop' ).addClass("disabled");
		$( '.shadowLine' ).addClass("disabled");
		$( '.shadowBottom' ).addClass("disabled");
	  }, function() {
		$( '.shadowTop2' ).addClass("disabled");
		$( '.shadowLine2' ).addClass("disabled");
		$( '.shadowBottom2' ).addClass("disabled");
		
		$( '.shadowTop' ).removeClass("disabled");
		$( '.shadowLine' ).removeClass("disabled");
		$( '.shadowBottom' ).removeClass("disabled");
	  }
	);
});

function showModalPopup(modalName, btnName) {
  const modal = document.getElementById(modalName);
  const btn = document.getElementById(btnName);

  if(!modal || !btn) return;

  const span = modal.getElementsByClassName("modal-close")[0];

  btn.onclick = function(e) {
    e.preventDefault()
    modal.style.display = "block";
  }

  span.onclick = function() {
    modal.style.display = "none";
  }

  window.onclick = function(event) {
    if (event.target == modal) {
      modal.style.display = "none";
    }
  };
}

function toggleGrooming() {
  $('body').toggleClass('grooming-page');
}

showModalPopup("addDocument", "addDocumentBtn")
showModalPopup("addPassport", "addPassportBtn")
showModalPopup("sendPayment", "sendPaymentBtn")

function initSidebar() {
  $('.sidebar-item').on('click', function(e) {
    e.preventDefault()
    if(!$(this).hasClass('active')) {
      $('.sidebar-item').removeClass('active');
      $(this).addClass('active');

      const srcUrl = $(this).data('src');

      $('.distribution-iframe').attr('src', srcUrl);
    }
  });

  $('.sidebar-mob').on('change', function() {
    const srcUrl = $(this).find(':selected').data('src');
    $('.distribution-iframe').attr('src', srcUrl);
  })
}

initSidebar()
