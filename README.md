# golang-crawling
당산포레스트 임대공고가 공지사항에 나오기에 주기적으로 확인 하기 위한 프로젝트

* golang사용
* config 정보을 file로 저장해서 관리
* quartz용 (asynq)
- 주기 하루 1번으로 생각중

https://www.포레나당산.com/notice.html

----
robots.txt

User-agent: *

Disallow:/board/adm/

Disallow:/prev/

Disallow:/test_pc/

Disallow:/test_mo/

Allow:/mobile/

Allow:/sitemap.xml

---
크롤링이 금지가 안됐음에 따라 공지사항을 크롤링 진행함.

제외가 추가되는 경우 해당 소스는 파기

크롤링 할 위치 https://www.포레나당산.com/board/bbs/board.php?bo_table=notice

---

고민중

인덱스 화면도 만들어서 file을 출력, 쿼츠 남은시간 출력
