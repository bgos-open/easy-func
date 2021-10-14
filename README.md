# easy-func

## 项目介绍
使用 golang 来翻译 php 函数。可以看做php函数在golang的映射字典。

## 项目要求

- 单元测试
    - 单测用例需参考 [https://github.com/php/php-src/tree/master/ext/standard/tests/strings](https://github.com/php/php-src/tree/master/ext/standard/tests/strings)
    - 单测覆盖率：go-carpet 或者 go-carpet --summary 保障单元测试覆盖率
    - 单测规范：goland -> general -> test for function 来生成单元测试用例代码
    - **单测覆盖率要求：100%**
- 圈复杂度
    - 在根目录使用命令 gocyclo -over 10 ./ 检查圈复杂度
    - **保证所有复杂度不超过 10**

## 函数
- [ ] [string](https://www.php.net/manual/en/ref.strings.php)
  - [ ] [addcslashes()](https://www.php.net/manual/en/function.addcslashes)
  - [ ] [addslashes()](https://www.php.net/manual/en/function.addslashes)
  - [ ] [bin2hex()](https://www.php.net/manual/en/function.bin2hex)
  - [ ] [chop()](https://www.php.net/manual/en/function.chop)
  - [x] [chr()](https://www.php.net/manual/en/function.chr)
  - [x] [chunk_split()](https://www.php.net/manual/en/function.chunk-split)
  - [ ] [~~convert_cyr_string()~~](https://www.php.net/manual/en/function.~~convert-cyr-string)
  - [ ] [convert_uudecode()](https://www.php.net/manual/en/function.convert-uudecode)
  - [ ] [convert_uuencode()](https://www.php.net/manual/en/function.convert-uuencode)
  - [ ] [count_chars()](https://www.php.net/manual/en/function.count-chars)
  - [x] [crc32()](https://www.php.net/manual/en/function.crc32)
  - [ ] [crypt()](https://www.php.net/manual/en/function.crypt)
  - [ ] [echo()](https://www.php.net/manual/en/function.echo)
  - [ ] [explode()](https://www.php.net/manual/en/function.explode)
  - [ ] [fprintf()](https://www.php.net/manual/en/function.fprintf)
  - [ ] [get_html_translation_table()](https://www.php.net/manual/en/function.get-html-translation-table)
  - [ ] [~~hebrev()~~](https://www.php.net/manual/en/function.~~hebrev)
  - [ ] [hebrevc()](https://www.php.net/manual/en/function.hebrevc)
  - [ ] [hex2bin()](https://www.php.net/manual/en/function.hex2bin)
  - [ ] [html_entity_decode()](https://www.php.net/manual/en/function.html-entity-decode)
  - [ ] [htmlentities()](https://www.php.net/manual/en/function.htmlentities)
  - [ ] [htmlspecialchars_decode()](https://www.php.net/manual/en/function.htmlspecialchars-decode)
  - [ ] [htmlspecialchars()](https://www.php.net/manual/en/function.htmlspecialchars)
  - [ ] [implode()](https://www.php.net/manual/en/function.implode)
  - [ ] [join()](https://www.php.net/manual/en/function.join)
  - [ ] [lcfirst()](https://www.php.net/manual/en/function.lcfirst)
  - [ ] [levenshtein()](https://www.php.net/manual/en/function.levenshtein)
  - [ ] [localeconv()](https://www.php.net/manual/en/function.localeconv)
  - [x] [ltrim()](https://www.php.net/manual/en/function.ltrim)
  - [ ] [md5_file()](https://www.php.net/manual/en/function.md5-file)
  - [x] [md5()](https://www.php.net/manual/en/function.md5)
  - [ ] [metaphone()](https://www.php.net/manual/en/function.metaphone)
  - [ ] [money_format()](https://www.php.net/manual/en/function.money-format)
  - [ ] [nl_langinfo()](https://www.php.net/manual/en/function.nl-langinfo)
  - [x] [nl2br()](https://www.php.net/manual/en/function.nl2br)
  - [ ] [number_format()](https://www.php.net/manual/en/function.number-format)
  - [x] [ord()](https://www.php.net/manual/en/function.ord)
  - [ ] [parse_str()](https://www.php.net/manual/en/function.parse-str)
  - [ ] [print()](https://www.php.net/manual/en/function.print)
  - [ ] [printf()](https://www.php.net/manual/en/function.printf)
  - [ ] [quoted_printable_decode()](https://www.php.net/manual/en/function.quoted-printable-decode)
  - [ ] [quoted_printable_encode()](https://www.php.net/manual/en/function.quoted-printable-encode)
  - [ ] [quotemeta()](https://www.php.net/manual/en/function.quotemeta)
  - [x] [rtrim()](https://www.php.net/manual/en/function.rtrim)
  - [ ] [setlocale()](https://www.php.net/manual/en/function.setlocale)
  - [ ] [sha1_file()](https://www.php.net/manual/en/function.sha1-file)
  - [ ] [sha1()](https://www.php.net/manual/en/function.sha1)
  - [ ] [similar_text()](https://www.php.net/manual/en/function.similar-text)
  - [ ] [soundex()](https://www.php.net/manual/en/function.soundex)
  - [ ] [sprintf()](https://www.php.net/manual/en/function.sprintf)
  - [ ] [sscanf()](https://www.php.net/manual/en/function.sscanf)
  - [ ] [str_contains()](https://www.php.net/manual/en/function.str-contains)
  - [ ] [str_ends_with()](https://www.php.net/manual/en/function.str-ends-with)
  - [ ] [str_getcsv()](https://www.php.net/manual/en/function.str-getcsv)
  - [ ] [str_ireplace()](https://www.php.net/manual/en/function.str-ireplace)
  - [x] [str_pad()](https://www.php.net/manual/en/function.str-pad)
  - [x] [str_repeat()](https://www.php.net/manual/en/function.str-repeat)
  - [ ] [str_replace()](https://www.php.net/manual/en/function.str-replace)
  - [ ] [str_rot13()](https://www.php.net/manual/en/function.str-rot13)
  - [ ] [str_shuffle()](https://www.php.net/manual/en/function.str-shuffle)
  - [ ] [str_split()](https://www.php.net/manual/en/function.str-split)
  - [ ] [str_starts_with()](https://www.php.net/manual/en/function.str-starts-with)
  - [ ] [str_word_count()](https://www.php.net/manual/en/function.str-word-count)
  - [ ] [strcasecmp()](https://www.php.net/manual/en/function.strcasecmp)
  - [ ] [strchr()](https://www.php.net/manual/en/function.strchr)
  - [ ] [strcmp()](https://www.php.net/manual/en/function.strcmp)
  - [ ] [strcoll()](https://www.php.net/manual/en/function.strcoll)
  - [ ] [strcspn()](https://www.php.net/manual/en/function.strcspn)
  - [ ] [strip_tags()](https://www.php.net/manual/en/function.strip-tags)
  - [ ] [stripcslashes()](https://www.php.net/manual/en/function.stripcslashes)
  - [ ] [stripos()](https://www.php.net/manual/en/function.stripos)
  - [ ] [stripslashes()](https://www.php.net/manual/en/function.stripslashes)
  - [ ] [stristr()](https://www.php.net/manual/en/function.stristr)
  - [ ] [strlen()](https://www.php.net/manual/en/function.strlen)
  - [ ] [strnatcasecmp()](https://www.php.net/manual/en/function.strnatcasecmp)
  - [ ] [strnatcmp()](https://www.php.net/manual/en/function.strnatcmp)
  - [ ] [strncasecmp()](https://www.php.net/manual/en/function.strncasecmp)
  - [ ] [strncmp()](https://www.php.net/manual/en/function.strncmp)
  - [ ] [strpbrk()](https://www.php.net/manual/en/function.strpbrk)
  - [ ] [strpos()](https://www.php.net/manual/en/function.strpos)
  - [ ] [strrchr()](https://www.php.net/manual/en/function.strrchr)
  - [ ] [strrev()](https://www.php.net/manual/en/function.strrev)
  - [ ] [strripos()](https://www.php.net/manual/en/function.strripos)
  - [ ] [strrpos()](https://www.php.net/manual/en/function.strrpos)
  - [ ] [strspn()](https://www.php.net/manual/en/function.strspn)
  - [x] [strstr()](https://www.php.net/manual/en/function.strstr)
  - [ ] [strtok()](https://www.php.net/manual/en/function.strtok)
  - [ ] [strtolower()](https://www.php.net/manual/en/function.strtolower)
  - [ ] [strtoupper()](https://www.php.net/manual/en/function.strtoupper)
  - [ ] [strtr()](https://www.php.net/manual/en/function.strtr)
  - [ ] [substr_compare()](https://www.php.net/manual/en/function.substr-compare)
  - [ ] [substr_count()](https://www.php.net/manual/en/function.substr-count)
  - [ ] [substr_replace()](https://www.php.net/manual/en/function.substr-replace)
  - [ ] [substr()](https://www.php.net/manual/en/function.substr)
  - [x] [trim()](https://www.php.net/manual/en/function.trim)
  - [ ] [ucfirst()](https://www.php.net/manual/en/function.ucfirst)
  - [x] [ucwords()](https://www.php.net/manual/en/function.ucwords)
  - [ ] [vfprintf()](https://www.php.net/manual/en/function.vfprintf)
  - [ ] [vprintf()](https://www.php.net/manual/en/function.vprintf)
  - [ ] [vsprintf()](https://www.php.net/manual/en/function.vsprintf)
  - [ ] [wordwrap()](https://www.php.net/manual/en/function.wordwrap)
- [ ] array

## 圈复杂度：

- 安装参考：https://github.com/fzipp/gocyclo
- 示例命令：

    `gocyclo -over 10 ./` #查看圈复杂度超过 10 的 function 
    
## 单元测试：

生成测试用例：goland 在方法名处右键，Generate, Tests for package. 根据生成的代码完善用例
示例命令：(在文件所在的文件夹下)：

 - 方法一：`go test` 
 - 方法二：`go test --cover` 
 - 方法三：在相应的 _test.go 文件下，用例前面有一个三角形，直接点击运行

## 单元测试覆盖率
https://github.com/msoap/go-carpet

