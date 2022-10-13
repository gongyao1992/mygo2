<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: simple.proto

namespace Pb;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>pb.SimpleRequest</code>
 */
class SimpleRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
     *
     * Generated from protobuf field <code>string data = 1;</code>
     */
    protected $data = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $data
     *           定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Simple::initOnce();
        parent::__construct($data);
    }

    /**
     * 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
     *
     * Generated from protobuf field <code>string data = 1;</code>
     * @return string
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
     *
     * Generated from protobuf field <code>string data = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setData($var)
    {
        GPBUtil::checkString($var, True);
        $this->data = $var;

        return $this;
    }

}

