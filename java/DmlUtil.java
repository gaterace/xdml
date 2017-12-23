package dml;

import com.google.protobuf.ByteString;
import com.google.protobuf.StringValue;

import java.util.UUID;

public class DmlUtil
{

    public static dml.DmlExtension.DateTime ConvertDateTime(java.sql.Timestamp date)
    {
        dml.DmlExtension.DateTime.Builder dtb = dml.DmlExtension.DateTime.newBuilder();
        
        if (date != null)
        {
            dtb.setMilliseconds(date.getTime());
        }
        
        return dtb.build();
    }
    
    public static java.sql.Timestamp ConvertDateTime(dml.DmlExtension.DateTime date)
    {
        java.sql.Timestamp result = null;
        if (date != null)
        {
            result = new java.sql.Timestamp(date.getMilliseconds());            
        }        

        return result;
    }

    public static java.sql.Date ConvertDate(dml.DmlExtension.DateTime date)
    {
        java.sql.Date result = null;
        if(date != null)
        {
            result = new java.sql.Date(date.getMilliseconds());
        }

        return result;
    }

    public static dml.DmlExtension.Decimal ConvertDecimal(java.math.BigDecimal decValue)
    {
        dml.DmlExtension.Decimal.Builder decb = dml.DmlExtension.Decimal.newBuilder();
        
        if (decValue != null)
        {
            decb.setPlaintext(decValue.toPlainString());
        }
        
        
        return decb.build();       
    }
    
    public static java.math.BigDecimal ConvertDecimal(dml.DmlExtension.Decimal decValue)
    {
        java.math.BigDecimal result = null;
        
        if (decValue != null)
        {
            result = new java.math.BigDecimal(decValue.getPlaintext());
        }        
        
        return result;
    }

    public static byte[] ConvertGuid(dml.DmlExtension.Guid guidVal) {
        byte[] result = null;

        if (guidVal != null) {
            result = guidVal.getGuid().toByteArray();
        }

        return result;
    }

    public static byte[] ConvertUuid(UUID uuidVal) {
        byte[] result = null;

        if (uuidVal != null) {
            String strVal = uuidVal.toString();
            StringBuffer sb = new StringBuffer();
            sb.append(strVal.substring(0,8));
            sb.append(strVal.substring(9,13));
            sb.append(strVal.substring(14,18));
            sb.append(strVal.substring(19,23));
            sb.append(strVal.substring(24));

            result = toByte(sb.toString());
        }

        return result;
    }

    public static UUID ConvertGuidToUuid(dml.DmlExtension.Guid guidVal) {
        UUID result = null;

        if (guidVal != null) {
            byte[] uuidBytes = guidVal.getGuid().toByteArray();
            String uuidString = ConvertUuid(uuidBytes);
            result = UUID.fromString(uuidString);
        }

        return result;
    }

    public static String ConvertUuid(byte[] uuidVal) {
        String result = null;

        if ((uuidVal != null) && (uuidVal.length == 16)) {
            String hexString = toHex(uuidVal);
            StringBuffer sb = new StringBuffer();
            sb.append(hexString.substring(0, 8));
            sb.append('-');
            sb.append(hexString.substring(8, 12));
            sb.append('-');
            sb.append(hexString.substring(12, 16));
            sb.append('-');
            sb.append(hexString.substring(16, 20));
            sb.append('-');
            sb.append(hexString.substring(20));
            result = sb.toString();
        }

        return result;
    }

    public static dml.DmlExtension.Guid ConvertGuid(byte[] guidVal) {
        dml.DmlExtension.Guid.Builder guid = dml.DmlExtension.Guid.newBuilder();

        if ((guidVal != null) && (guidVal.length == 16)) {
            guid.setGuid(ByteString.copyFrom(guidVal));
        }

        return guid.build();
    }


    public static dml.DmlExtension.Guid ConvertGuid(UUID uuidVal) {
        dml.DmlExtension.Guid.Builder guid = dml.DmlExtension.Guid.newBuilder();

        if (uuidVal != null) {
            byte[] bval = ConvertUuid(uuidVal);
            guid.setGuid(ByteString.copyFrom(bval));
        }

        return guid.build();
    }

    public static byte[] ConvertByteArray(com.google.protobuf.ByteString bval) {
        byte[] result = null;

        if (bval != null) {
            result = bval.toByteArray();
        }

        return result;
    }

    public static com.google.protobuf.ByteString ConvertByteArray(byte[] bval) {
        com.google.protobuf.ByteString result = null;

        if (bval != null) {
            result = com.google.protobuf.ByteString.copyFrom(bval);
        }

        return result;
    }

    public static byte[] toByte(String hexString) {
        int len = hexString.length()/2;
        byte[] result = new byte[len];
        for (int i = 0; i < len; i++)
            result[i] = Integer.valueOf(hexString.substring(2*i, 2*i+2), 16).byteValue();
        return result;
    }

    public static String toHex(byte[] buf) {
        if (buf == null)
            return "";
        StringBuffer result = new StringBuffer(2*buf.length);
        for (int i = 0; i < buf.length; i++) {
            appendHex(result, buf[i]);
        }
        return result.toString();
    }
    private final static String HEX = "0123456789ABCDEF";
    private static void appendHex(StringBuffer sb, byte b) {
        sb.append(HEX.charAt((b>>4)&0x0f)).append(HEX.charAt(b&0x0f));
    }


}
