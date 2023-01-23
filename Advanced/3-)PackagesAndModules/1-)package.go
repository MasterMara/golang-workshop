package main

func main() {

}

/*
1-) Resources:
	-https://www.youtube.com/watch?v=nLaxs5w9bZc


GoPath: - GoPath, Go dilinin çalışma zamanında kullandığı dosya yolunu belirtir.
		GoPath, Go dilinin kullanımı sırasında hangi dizinlerde arama yapacağını ve hangi paketleri kullanacağını belirler.

GoRoot: - GOROOT, Go yazılım dilinin kurulu olduğu kök dizinin yolunu belirtir.
		  Go dilinin kurulu olduğu dizin, Go dilinin kurulu olduğu bilgisayarda Go dilinin standart kütüphane ve araçlarının bulunduğu dizindir
 		- GOROOT, Go dilinde bir sistem değişkenidir ve Go dilinin kurulu olduğu dizini gösterir

GoMod: GoMod, Go dilinde kullanılan bir paket yönetim aracıdır. Go dilinin kütüphanelerini ve bağımlılıklarını yönetmek için tasarlandı
	  -GoMod, bir projeye dahil edilen tüm paketlerin ve bağımlılıklarının sürümlerini izler ve projeyi derlerken gerekli olan tüm paketleri
	  - indirir. Bu, projeyi çalıştırmak için gerekli olan tüm paketleri
       tek bir yerde tutmanıza ve projeyi kolayca dağıtmak için gerekli olan tüm paketlerin bir listesini oluşturmanıza yardımcı olur.
	  - GoMod ayrıca, projenizdeki paketleri çeşitli kaynaklardan indirir. Bu kaynaklar arasında GitHub, Bitbucket ve Google Cloud Storage gibi
      popüler depolama hizmetleri bulunur. Bu, projenizdeki paketleri çeşitli kaynaklardan indirme ve güncelleme olanağı sağlar.

GoProxy:

GoSum: package dependency chain not our dependency but our dependency package's dependency in here. .It is generated file.


2-) Every go file has to be some part of package
3-) go mod init <name> will create go module
4-) go mod tidy will tidy unused dependency and also download missing dependecy
5-) go get <packageName> will download package
5-) go mod why <packageName> will say us to why we need that package
6-) go mod download (will dowloand all package inside go mod file)
6-) import denemeProject/util (moduleName/packageName) if you want to import someThing other
7-) we can use variables in the same package even if they are in different files
8-) if we are not using that dependecy go.mod will mark as a //indirect
9-) import v3 "github/something" şeklinde v3. diyerek o paketin içerisine erişebilriiz.
*/
