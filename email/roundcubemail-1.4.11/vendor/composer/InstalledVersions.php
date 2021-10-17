<?php











namespace Composer;

use Composer\Autoload\ClassLoader;
use Composer\Semver\VersionParser;






class InstalledVersions
{
private static $installed = array (
  'root' => 
  array (
    'pretty_version' => 'dev-master',
    'version' => 'dev-master',
    'aliases' => 
    array (
    ),
    'reference' => '39da42bacb746d323e00ce5e293fc0463db7baf0',
    'name' => 'roundcube/roundcubemail',
  ),
  'versions' => 
  array (
    'endroid/qr-code' => 
    array (
      'pretty_version' => '1.6.6',
      'version' => '1.6.6.0',
      'aliases' => 
      array (
      ),
      'reference' => 'cef5d5b7b904d7bb0708eb744c35316364b65fa0',
    ),
    'kolab/net_ldap3' => 
    array (
      'pretty_version' => 'v1.1.3',
      'version' => '1.1.3.0',
      'aliases' => 
      array (
      ),
      'reference' => 'e1835e36fa5d3434f9804670fc3cdd16992ec66a',
    ),
    'masterminds/html5' => 
    array (
      'pretty_version' => '2.5.0',
      'version' => '2.5.0.0',
      'aliases' => 
      array (
      ),
      'reference' => 'b5d892a4bd058d61f736935d32a9c248f11ccc93',
    ),
    'pear/auth_sasl' => 
    array (
      'pretty_version' => 'v1.1.0',
      'version' => '1.1.0.0',
      'aliases' => 
      array (
      ),
      'reference' => 'db1ead3dc0bf986d2bab0dbc04d114800cf91dee',
    ),
    'pear/console_commandline' => 
    array (
      'pretty_version' => 'v1.2.3',
      'version' => '1.2.3.0',
      'aliases' => 
      array (
      ),
      'reference' => '331b295d00c9fa2dae30fe14538d430d72d954c6',
    ),
    'pear/console_getopt' => 
    array (
      'pretty_version' => 'v1.4.3',
      'version' => '1.4.3.0',
      'aliases' => 
      array (
      ),
      'reference' => 'a41f8d3e668987609178c7c4a9fe48fecac53fa0',
    ),
    'pear/crypt_gpg' => 
    array (
      'pretty_version' => 'v1.6.4',
      'version' => '1.6.4.0',
      'aliases' => 
      array (
      ),
      'reference' => '7aa6a2a2b0017acb7eb9adf109ee4bfee0bf1d2f',
    ),
    'pear/mail_mime' => 
    array (
      'pretty_version' => '1.10.9',
      'version' => '1.10.9.0',
      'aliases' => 
      array (
      ),
      'reference' => '1e7ae4e5258b6c0d385a8e76add567934245d38d',
    ),
    'pear/net_idna2' => 
    array (
      'pretty_version' => 'v0.2.0',
      'version' => '0.2.0.0',
      'aliases' => 
      array (
      ),
      'reference' => '51734eaf8be2df58e8aad5835b9966459b2fb37c',
    ),
    'pear/net_ldap2' => 
    array (
      'pretty_version' => 'v2.2.0',
      'version' => '2.2.0.0',
      'aliases' => 
      array (
      ),
      'reference' => '38f1b22a96dfbd7ec53852f0e1e7ec1a9a5eb0e8',
    ),
    'pear/net_sieve' => 
    array (
      'pretty_version' => '1.4.4',
      'version' => '1.4.4.0',
      'aliases' => 
      array (
      ),
      'reference' => '5bd8f14db4684f75aee468c0ff8c5ff5175f4716',
    ),
    'pear/net_smtp' => 
    array (
      'pretty_version' => '1.8.1',
      'version' => '1.8.1.0',
      'aliases' => 
      array (
      ),
      'reference' => '8f93b34878ef16844c754d1a1211c1e3677fa07f',
    ),
    'pear/net_socket' => 
    array (
      'pretty_version' => 'v1.2.2',
      'version' => '1.2.2.0',
      'aliases' => 
      array (
      ),
      'reference' => 'bbe6a12bb4f7059dba161f6ddd43f369c0ec8d09',
    ),
    'pear/pear-core-minimal' => 
    array (
      'pretty_version' => 'v1.10.10',
      'version' => '1.10.10.0',
      'aliases' => 
      array (
      ),
      'reference' => '625a3c429d9b2c1546438679074cac1b089116a7',
    ),
    'pear/pear_exception' => 
    array (
      'pretty_version' => 'v1.0.1',
      'version' => '1.0.1.0',
      'aliases' => 
      array (
      ),
      'reference' => 'dbb42a5a0e45f3adcf99babfb2a1ba77b8ac36a7',
    ),
    'roundcube/plugin-installer' => 
    array (
      'pretty_version' => '0.2.0',
      'version' => '0.2.0.0',
      'aliases' => 
      array (
      ),
      'reference' => '10148d050ee2db019d83113a479fb7716fd04d48',
    ),
    'roundcube/roundcubemail' => 
    array (
      'pretty_version' => 'dev-master',
      'version' => 'dev-master',
      'aliases' => 
      array (
      ),
      'reference' => '39da42bacb746d323e00ce5e293fc0463db7baf0',
    ),
    'rsky/pear-core-min' => 
    array (
      'replaced' => 
      array (
        0 => 'v1.10.10',
      ),
    ),
  ),
);
private static $canGetVendors;
private static $installedByVendor = array();







public static function getInstalledPackages()
{
$packages = array();
foreach (self::getInstalled() as $installed) {
$packages[] = array_keys($installed['versions']);
}


if (1 === \count($packages)) {
return $packages[0];
}

return array_keys(array_flip(\call_user_func_array('array_merge', $packages)));
}









public static function isInstalled($packageName)
{
foreach (self::getInstalled() as $installed) {
if (isset($installed['versions'][$packageName])) {
return true;
}
}

return false;
}














public static function satisfies(VersionParser $parser, $packageName, $constraint)
{
$constraint = $parser->parseConstraints($constraint);
$provided = $parser->parseConstraints(self::getVersionRanges($packageName));

return $provided->matches($constraint);
}










public static function getVersionRanges($packageName)
{
foreach (self::getInstalled() as $installed) {
if (!isset($installed['versions'][$packageName])) {
continue;
}

$ranges = array();
if (isset($installed['versions'][$packageName]['pretty_version'])) {
$ranges[] = $installed['versions'][$packageName]['pretty_version'];
}
if (array_key_exists('aliases', $installed['versions'][$packageName])) {
$ranges = array_merge($ranges, $installed['versions'][$packageName]['aliases']);
}
if (array_key_exists('replaced', $installed['versions'][$packageName])) {
$ranges = array_merge($ranges, $installed['versions'][$packageName]['replaced']);
}
if (array_key_exists('provided', $installed['versions'][$packageName])) {
$ranges = array_merge($ranges, $installed['versions'][$packageName]['provided']);
}

return implode(' || ', $ranges);
}

throw new \OutOfBoundsException('Package "' . $packageName . '" is not installed');
}





public static function getVersion($packageName)
{
foreach (self::getInstalled() as $installed) {
if (!isset($installed['versions'][$packageName])) {
continue;
}

if (!isset($installed['versions'][$packageName]['version'])) {
return null;
}

return $installed['versions'][$packageName]['version'];
}

throw new \OutOfBoundsException('Package "' . $packageName . '" is not installed');
}





public static function getPrettyVersion($packageName)
{
foreach (self::getInstalled() as $installed) {
if (!isset($installed['versions'][$packageName])) {
continue;
}

if (!isset($installed['versions'][$packageName]['pretty_version'])) {
return null;
}

return $installed['versions'][$packageName]['pretty_version'];
}

throw new \OutOfBoundsException('Package "' . $packageName . '" is not installed');
}





public static function getReference($packageName)
{
foreach (self::getInstalled() as $installed) {
if (!isset($installed['versions'][$packageName])) {
continue;
}

if (!isset($installed['versions'][$packageName]['reference'])) {
return null;
}

return $installed['versions'][$packageName]['reference'];
}

throw new \OutOfBoundsException('Package "' . $packageName . '" is not installed');
}





public static function getRootPackage()
{
$installed = self::getInstalled();

return $installed[0]['root'];
}







public static function getRawData()
{
return self::$installed;
}



















public static function reload($data)
{
self::$installed = $data;
self::$installedByVendor = array();
}




private static function getInstalled()
{
if (null === self::$canGetVendors) {
self::$canGetVendors = method_exists('Composer\Autoload\ClassLoader', 'getRegisteredLoaders');
}

$installed = array();

if (self::$canGetVendors) {

foreach (ClassLoader::getRegisteredLoaders() as $vendorDir => $loader) {
if (isset(self::$installedByVendor[$vendorDir])) {
$installed[] = self::$installedByVendor[$vendorDir];
} elseif (is_file($vendorDir.'/composer/installed.php')) {
$installed[] = self::$installedByVendor[$vendorDir] = require $vendorDir.'/composer/installed.php';
}
}
}

$installed[] = self::$installed;

return $installed;
}
}
